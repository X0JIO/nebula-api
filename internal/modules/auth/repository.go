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

// ----------------------------------------------------------------------
// Refresh Tokens
// ----------------------------------------------------------------------

func (r *Repository) CreateRefreshToken(
	ctx context.Context,
	userID string,
	hash string,
	expiresAt pgtype.Timestamp,
) (db.RefreshToken, error) {

	var uid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return db.RefreshToken{}, err
	}

	return r.queries.CreateRefreshToken(
		ctx,
		db.CreateRefreshTokenParams{
			UserID:    uid,
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

func (r *Repository) RevokeAllRefreshTokens(
	ctx context.Context,
	userID pgtype.UUID,
) error {
	return r.queries.RevokeAllRefreshTokens(
		ctx,
		userID,
	)
}

// ----------------------------------------------------------------------
// Sessions
// ----------------------------------------------------------------------

func (r *Repository) CreateSession(
	ctx context.Context,
	arg db.CreateSessionParams,
) (db.Session, error) {

	return r.queries.CreateSession(
		ctx,
		arg,
	)
}

func (r *Repository) GetSessionByRefreshHash(
	ctx context.Context,
	hash string,
) (db.Session, error) {

	return r.queries.GetSessionByRefreshHash(
		ctx,
		hash,
	)
}

func (r *Repository) UpdateSessionRefresh(
	ctx context.Context,
	arg db.UpdateSessionRefreshParams,
) error {

	return r.queries.UpdateSessionRefresh(
		ctx,
		arg,
	)
}

func (r *Repository) RevokeSession(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.queries.RevokeSession(
		ctx,
		id,
	)
}

func (r *Repository) RevokeUserSessions(
	ctx context.Context,
	userID pgtype.UUID,
) error {

	return r.queries.RevokeUserSessions(
		ctx,
		userID,
	)
}

func (r *Repository) ListSessions(
	ctx context.Context,
	userID pgtype.UUID,
) ([]db.Session, error) {

	return r.queries.ListSessions(
		ctx,
		userID,
	)
}

func (r *Repository) UpdateRefreshToken(
	ctx context.Context,
	oldHash string,
	newHash string,
	expires pgtype.Timestamp,
) error {

	return r.queries.UpdateRefreshToken(
		ctx,
		db.UpdateRefreshTokenParams{
			TokenHash:   oldHash,
			TokenHash_2: newHash,
			ExpiresAt:   expires,
		},
	)
}
