package server

import (
	"context"
	"fmt"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	params db.CreateVPNServerParams,
) (db.VpnServer, error) {
	return r.q.CreateVPNServer(ctx, params)
}

func (r *Repository) List(
	ctx context.Context,
) ([]db.VpnServer, error) {
	return r.q.ListVPNServers(ctx)
}

func (r *Repository) Get(ctx context.Context, id pgtype.UUID) (db.VpnServer, error) {
	return r.q.GetVPNServer(ctx, id)
}

func (r *Repository) GetActive(
	ctx context.Context,
) (db.VpnServer, error) {
	return r.q.GetActiveVPNServer(ctx)
}

func (r *Repository) Deactivate(
	ctx context.Context,
	id pgtype.UUID,
) (db.VpnServer, error) {
	return r.q.DeactivateVPNServer(ctx, id)
}

func (r *Repository) Update(
	ctx context.Context,
	params db.UpdateVPNServerParams,
) (db.VpnServer, error) {
	return r.q.UpdateVPNServer(ctx, params)
}

func (r *Repository) Delete(ctx context.Context, id pgtype.UUID) error {
	return r.q.DeleteVPNServer(ctx, id)
}

func (r *Repository) DeactivateAll(
	ctx context.Context,
) error {
	return r.q.DeactivateAllVPNServers(ctx)
}

func (r *Repository) Activate(ctx context.Context, id pgtype.UUID) (db.VpnServer, error) {
	return r.q.ActivateVPNServer(ctx, id)
}

func toUUID(id interface{}) (pgtype.UUID, error) {
	switch v := id.(type) {
	case pgtype.UUID:
		return v, nil

	case string:
		var uuid pgtype.UUID
		if err := uuid.Scan(v); err != nil {
			return pgtype.UUID{}, err
		}
		return uuid, nil

	default:
		return pgtype.UUID{}, fmt.Errorf("invalid UUID type %T", id)
	}
}
