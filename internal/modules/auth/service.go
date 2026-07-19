package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
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

type RefreshRepository interface {
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
}

type Service struct {
	users   UserRepository
	refresh RefreshRepository
	jwt     *JWT
	cfg     config.JWTConfig
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewService(
	users UserRepository,
	refresh RefreshRepository,
	jwt *JWT,
	cfg config.JWTConfig,
) *Service {

	return &Service{
		users:   users,
		refresh: refresh,
		jwt:     jwt,
		cfg:     cfg,
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

	return s.users.Create(ctx, email, hash)
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

	user, err := s.users.GetByEmail(ctx, email)
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
) (Tokens, error) {

	user, err := s.Login(ctx, email, password)
	if err != nil {
		return Tokens{}, err
	}

	access, err := s.jwt.GenerateAccessToken(
		user.ID.String(),
		user.Role,
		s.cfg.AccessTTL,
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

	hash := sha256.Sum256([]byte(refreshToken))
	hashString := hex.EncodeToString(hash[:])

	expires := pgtype.Timestamp{
		Time:  time.Now().Add(s.cfg.RefreshTTL),
		Valid: true,
	}

	_, err = s.refresh.CreateRefreshToken(
		ctx,
		user.ID.String(),
		hashString,
		expires,
	)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  access,
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

	userID := claims.Subject

	hash := sha256.Sum256([]byte(refreshToken))
	hashString := hex.EncodeToString(hash[:])

	token, err := s.refresh.GetRefreshToken(ctx, hashString)
	if err != nil {
		return Tokens{}, errors.New("refresh token not found")
	}

	if !token.ExpiresAt.Valid || token.ExpiresAt.Time.Before(time.Now()) {
		return Tokens{}, errors.New("refresh token expired")
	}

	if err := s.refresh.RevokeRefreshToken(ctx, hashString); err != nil {
		return Tokens{}, err
	}

	user, err := s.users.GetByID(ctx, userID)
	if err != nil {
		return Tokens{}, err
	}

	newAccess, err := s.jwt.GenerateAccessToken(
		user.ID.String(),
		user.Role,
		s.cfg.AccessTTL,
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

	newHash := sha256.Sum256([]byte(newRefresh))
	newHashString := hex.EncodeToString(newHash[:])

	expires := pgtype.Timestamp{
		Time:  time.Now().Add(s.cfg.RefreshTTL),
		Valid: true,
	}

	_, err = s.refresh.CreateRefreshToken(
		ctx,
		user.ID.String(),
		newHashString,
		expires,
	)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}, nil
}
