package auth

import (
	"context"
	"errors"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

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
}

func NewService(
	users ServiceRepository,
) *Service {
	return &Service{
		users: users,
	}
}

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
