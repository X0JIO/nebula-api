package auth

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(
	queries *db.Queries,
) *Repository {

	return &Repository{
		queries: queries,
	}
}

func (r *Repository) CreateRefreshToken(
	ctx context.Context,
	userID string,
	hash string,
	expiresAt pgtype.Timestamp,
) (db.RefreshToken, error) {

	var id pgtype.UUID

	err := id.Scan(userID)
	if err != nil {
		return db.RefreshToken{}, err
	}

	return r.queries.CreateRefreshToken(
		ctx,
		db.CreateRefreshTokenParams{
			UserID:    id,
			TokenHash: hash,
			ExpiresAt: expiresAt,
		},
	)
}

func (r *Repository) GetRefreshToken(
	ctx context.Context,
	hash string,
) (db.RefreshToken, error) {

	return r.queries.GetRefreshToken(
		ctx,
		hash,
	)
}

func (r *Repository) RevokeRefreshToken(
	ctx context.Context,
	hash string,
) error {

	return r.queries.RevokeRefreshToken(
		ctx,
		hash,
	)
}
