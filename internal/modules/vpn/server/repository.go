package server

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(
	q *db.Queries,
) *Repository {

	return &Repository{
		q: q,
	}
}

func (r *Repository) GetActive(
	ctx context.Context,
) (db.VpnServer, error) {

	return r.q.GetActiveVPNServer(ctx)

}
