package users

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

type Service struct {
	repository *Repository
}

func NewService(
	repository *Repository,
) *Service {

	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(
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

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return db.User{}, err
	}

	return s.repository.Create(
		ctx,
		email,
		string(hash),
	)
}
