package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"net/netip"
	"strings"
	"time"

	"github.com/X0JIO/nebula-api/internal/platform/config"
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface {
	GetByID(
		context.Context,
		string,
	) (db.User, error)

	GetByEmail(
		context.Context,
		string,
	) (db.User, error)

	Create(
		context.Context,
		string,
		string,
	) (db.User, error)
}

type AuthRepository interface {

	// refresh_tokens

	CreateRefreshToken(
		context.Context,
		string,
		string,
		pgtype.Timestamp,
	) (db.RefreshToken, error)

	GetRefreshToken(
		context.Context,
		string,
	) (db.RefreshToken, error)

	RevokeRefreshToken(
		context.Context,
		string,
	) error

	RevokeAllRefreshTokens(
		context.Context,
		pgtype.UUID,
	) error

	UpdateRefreshToken(
		context.Context,
		string,
		string,
		pgtype.Timestamp,
	) error

	RevokeSession(
		context.Context,
		pgtype.UUID,
	) error

	RevokeUserSessions(
		context.Context,
		pgtype.UUID,
	) error

	// sessions

	CreateSession(
		context.Context,
		db.CreateSessionParams,
	) (db.Session, error)

	GetSessionByRefreshHash(
		context.Context,
		string,
	) (db.Session, error)

	UpdateSessionRefresh(
		context.Context,
		db.UpdateSessionRefreshParams,
	) error
}

type Service struct {
	users   UserRepository
	auth    AuthRepository
	devices DeviceService

	jwt *JWT
	cfg config.JWTConfig
}

type DeviceService interface {
	Register(
		ctx context.Context,
		userID pgtype.UUID,
		sessionID pgtype.UUID,
		deviceName string,
		platform string,
		fingerprint string,
		vpnUUID pgtype.UUID,
		ip string,
	) error
}

func NewService(
	users UserRepository,
	auth AuthRepository,
	devices DeviceService,
	jwt *JWT,
	cfg config.JWTConfig,
) *Service {

	return &Service{
		users:   users,
		auth:    auth,
		devices: devices,
		jwt:     jwt,
		cfg:     cfg,
	}
}

func DetectPlatform(ua string) string {
	switch {
	case strings.Contains(ua, "Windows"):
		return "Windows"
	case strings.Contains(ua, "Android"):
		return "Android"
	case strings.Contains(ua, "iPhone"):
		return "iOS"
	case strings.Contains(ua, "Macintosh"):
		return "macOS"
	case strings.Contains(ua, "Linux"):
		return "Linux"
	default:
		return "Unknown"
	}
}

func (s *Service) Register(
	ctx context.Context,
	email string,
	password string,
) (db.User, error) {

	if email == "" {
		return db.User{}, errors.New("email required")
	}

	if password == "" {
		return db.User{}, errors.New("password required")
	}

	hash, err := HashPassword(password)
	if err != nil {
		return db.User{}, err
	}

	return s.users.Create(
		ctx,
		email,
		hash,
	)
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (db.User, error) {

	if email == "" {
		return db.User{}, errors.New("email required")
	}

	if password == "" {
		return db.User{}, errors.New("password required")
	}

	user, err := s.users.GetByEmail(
		ctx,
		email,
	)
	if err != nil {
		return db.User{}, err
	}

	if err := CheckPassword(
		user.PasswordHash,
		password,
	); err != nil {
		return db.User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *Service) LoginTokens(
	ctx context.Context,
	email string,
	password string,
	deviceName string,
	ip string,
	userAgent string,
	fingerprint string,
) (Tokens, error) {

	user, err := s.Login(
		ctx,
		email,
		password,
	)

	if err != nil {
		return Tokens{}, err
	}

	refreshToken, err := s.jwt.GenerateRefreshToken(
		user.ID.String(),
		s.cfg.RefreshTTL,
	)
	if err != nil {
		return Tokens{}, err
	}

	hash := sha256.Sum256(
		[]byte(refreshToken),
	)

	hashString := hex.EncodeToString(
		hash[:],
	)

	expires := pgtype.Timestamp{
		Time:  time.Now().Add(s.cfg.RefreshTTL),
		Valid: true,
	}

	_, err = s.auth.CreateRefreshToken(
		ctx,
		user.ID.String(),
		hashString,
		expires,
	)
	if err != nil {
		return Tokens{}, err
	}

	var uid pgtype.UUID

	if err := uid.Scan(user.ID.String()); err != nil {
		return Tokens{}, err
	}

	addr, err := netip.ParseAddr(ip)
	if err != nil {
		addr = netip.IPv4Unspecified()
	}

	session, err := s.auth.CreateSession(
		ctx,
		db.CreateSessionParams{
			UserID: uid,

			RefreshTokenHash: hashString,

			DeviceName: pgtype.Text{
				String: deviceName,
				Valid:  deviceName != "",
			},

			Ip: &addr,

			UserAgent: pgtype.Text{
				String: userAgent,
				Valid:  userAgent != "",
			},

			ExpiresAt: pgtype.Timestamptz{
				Time:  expires.Time,
				Valid: true,
			},
		},
	)
	if err != nil {
		return Tokens{}, err
	}
	log.Printf(
		"LoginTokens: user=%s fingerprint=%q device=%q session=%s",
		user.ID.String(),
		fingerprint,
		deviceName,
		session.ID.String(),
	)

	accessToken, err := s.jwt.GenerateAccessToken(
		user.ID.String(),
		user.Role,
		session.ID.String(),
		s.cfg.AccessTTL,
	)
	if err != nil {
		return Tokens{}, err
	}

	platform := DetectPlatform(userAgent)

	err = s.devices.Register(
		ctx,
		user.ID,
		session.ID,
		deviceName,
		platform,
		fingerprint,
		pgtype.UUID{}, // NULL
		ip,
	)
	if err != nil {
		log.Printf("devices.Register error: %v", err)
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) Refresh(
	ctx context.Context,
	refreshToken string,
) (Tokens, error) {

	claims, err := s.jwt.ParseToken(refreshToken)
	if err != nil {
		return Tokens{}, errors.New("invalid refresh token")
	}

	if claims.Type != "refresh" {
		return Tokens{}, errors.New("invalid token type")
	}

	hash := sha256.Sum256([]byte(refreshToken))
	hashString := hex.EncodeToString(hash[:])

	token, err := s.auth.GetRefreshToken(
		ctx,
		hashString,
	)
	if err != nil {
		return Tokens{}, errors.New("refresh token not found")
	}

	session, err := s.auth.GetSessionByRefreshHash(
		ctx,
		hashString,
	)
	if err != nil {
		return Tokens{}, err
	}

	if !token.ExpiresAt.Valid || token.ExpiresAt.Time.Before(time.Now()) {
		return Tokens{}, errors.New("refresh token expired")
	}

	user, err := s.users.GetByID(
		ctx,
		claims.Subject,
	)
	if err != nil {
		return Tokens{}, err
	}

	newRefresh, err := s.jwt.GenerateRefreshToken(
		user.ID.String(),
		s.cfg.RefreshTTL,
	)
	if err != nil {
		return Tokens{}, err
	}

	newHashBytes := sha256.Sum256([]byte(newRefresh))
	newHash := hex.EncodeToString(newHashBytes[:])

	expires := pgtype.Timestamp{
		Time:  time.Now().Add(s.cfg.RefreshTTL),
		Valid: true,
	}

	err = s.auth.UpdateSessionRefresh(
		ctx,
		db.UpdateSessionRefreshParams{
			ID:               session.ID,
			RefreshTokenHash: newHash,
			ExpiresAt: pgtype.Timestamptz{
				Time:  expires.Time,
				Valid: true,
			},
		},
	)
	if err != nil {
		return Tokens{}, err
	}

	err = s.auth.UpdateRefreshToken(
		ctx,
		hashString,
		newHash,
		expires,
	)
	if err != nil {
		return Tokens{}, err
	}

	accessToken, err := s.jwt.GenerateAccessToken(
		user.ID.String(),
		user.Role,
		session.ID.String(),
		s.cfg.AccessTTL,
	)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: newRefresh,
	}, nil
}

func (s *Service) Logout(
	ctx context.Context,
	sessionID string,
	refreshHash string,
) error {

	var sid pgtype.UUID

	if err := sid.Scan(sessionID); err != nil {
		return err
	}

	if err := s.auth.RevokeSession(ctx, sid); err != nil {
		return err
	}

	return s.auth.RevokeRefreshToken(ctx, refreshHash)
}

func (s *Service) LogoutAll(
	ctx context.Context,
	userID string,
) error {

	var uid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return err
	}

	if err := s.auth.RevokeUserSessions(ctx, uid); err != nil {
		return err
	}

	return s.auth.RevokeAllRefreshTokens(
		ctx,
		uid,
	)
}
