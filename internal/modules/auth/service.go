package auth

import (
	"context"
	"errors"
	"time"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

type ServiceRepository interface {
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
type Service struct {
	users interface {
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

	jwt *JWT
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewService(
	users ServiceRepository,
	jwt *JWT,
) *Service {
	return &Service{
		users: users,
		jwt:   jwt,
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
) (Tokens, error) {

	user, err := s.Login(
		ctx,
		email,
		password,
	)

	if err != nil {
		return Tokens{}, err
	}

	access, err := s.jwt.GenerateAccessToken(
		user.ID.String(),
		15*time.Minute,
	)

	if err != nil {
		return Tokens{}, err
	}

	refresh, err := s.jwt.GenerateRefreshToken(
		user.ID.String(),
		720*time.Hour,
	)

	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *Service) Refresh(
	refreshToken string,
) (Tokens, error) {

	userID, err := s.jwt.ParseToken(
		refreshToken,
	)

	if err != nil {
		return Tokens{}, err
	}

	access, err := s.jwt.GenerateAccessToken(
		userID,
		15*time.Minute,
	)

	if err != nil {
		return Tokens{}, err
	}

	refresh, err := s.jwt.GenerateRefreshToken(
		userID,
		720*time.Hour,
	)

	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
