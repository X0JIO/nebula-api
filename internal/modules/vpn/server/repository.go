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

func (r *Repository) Create(
	ctx context.Context,
	params db.CreateVPNServerParams,
) (db.VpnServer, error) {

	return r.q.CreateVPNServer(
		ctx,
		params,
	)
}

func (r *Repository) List(
	ctx context.Context,
) ([]db.VpnServer, error) {

	return r.q.ListVPNServers(ctx)
}

func (r *Repository) GetActive(
	ctx context.Context,
) (db.VpnServer, error) {

	return r.q.GetActiveVPNServer(ctx)
}
