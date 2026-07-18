package users

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) *Repository {
	return &Repository{
		queries: queries,
	}
}

func (r *Repository) GetByID(
	ctx context.Context,
	id string,
) (db.User, error) {

	userID := pgtype.UUID{}

	err := userID.Scan(id)
	if err != nil {
		return db.User{}, err
	}

	return r.queries.GetUserByID(ctx, userID)
}

func (r *Repository) GetByEmail(
	ctx context.Context,
	email string,
) (db.User, error) {

	return r.queries.GetUserByEmail(ctx, email)
}

func (r *Repository) Create(
	ctx context.Context,
	email string,
	passwordHash string,
) (db.User, error) {

	return r.queries.CreateUser(
		ctx,
		db.CreateUserParams{
			Email:        email,
			PasswordHash: passwordHash,
		},
	)
}
