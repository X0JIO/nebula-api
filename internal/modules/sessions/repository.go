package sessions

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q db.Querier
}

func NewRepository(q db.Querier) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	arg db.CreateSessionParams,
) (db.Session, error) {

	return r.q.CreateSession(ctx, arg)
}

func (r *Repository) ByRefreshHash(
	ctx context.Context,
	hash string,
) (db.Session, error) {

	return r.q.GetSessionByRefreshHash(ctx, hash)
}

func (r *Repository) List(
	ctx context.Context,
	userID pgtype.UUID,
) ([]db.Session, error) {

	return r.q.ListSessions(ctx, userID)
}

func (r *Repository) Revoke(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.q.RevokeSession(ctx, id)
}

func (r *Repository) RevokeAll(
	ctx context.Context,
	userID pgtype.UUID,
) error {

	return r.q.RevokeUserSessions(ctx, userID)
}

func (r *Repository) RotateRefresh(
	ctx context.Context,
	arg db.UpdateSessionRefreshParams,
) error {

	return r.q.UpdateSessionRefresh(ctx, arg)
}

func (r *Repository) Delete(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.q.DeleteSession(
		ctx,
		id,
	)
}

func (r *Repository) Get(
	ctx context.Context,
	id pgtype.UUID,
) (db.Session, error) {

	return r.q.GetSession(ctx, id)
}

func (r *Repository) GetByID(
	ctx context.Context,
	id string,
) (db.Session, error) {

	var sid pgtype.UUID

	if err := sid.Scan(id); err != nil {
		return db.Session{}, err
	}

	return r.q.GetSession(ctx, sid)
}
