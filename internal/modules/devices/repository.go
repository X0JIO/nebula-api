package devices

import (
	"context"
	"net/netip"

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
	arg db.CreateDeviceParams,
) (db.Device, error) {

	return r.q.CreateDevice(ctx, arg)
}

func (r *Repository) Get(
	ctx context.Context,
	id pgtype.UUID,
) (db.Device, error) {

	return r.q.GetDevice(ctx, id)
}

func (r *Repository) List(
	ctx context.Context,
	userID pgtype.UUID,
) ([]db.Device, error) {

	return r.q.ListDevices(ctx, userID)
}

func (r *Repository) UpdateLastSeen(
	ctx context.Context,
	arg db.UpdateDeviceLastSeenParams,
) error {

	return r.q.UpdateDeviceLastSeen(ctx, arg)
}

func (r *Repository) Delete(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.q.DeleteDevice(ctx, id)
}

func (r *Repository) DeleteUserDevices(
	ctx context.Context,
	userID pgtype.UUID,
) error {

	return r.q.DeleteUserDevices(ctx, userID)
}

func (r *Repository) GetByFingerprint(
	ctx context.Context,
	userID pgtype.UUID,
	fingerprint string,
) (db.Device, error) {

	return r.q.GetDeviceByFingerprint(
		ctx,
		db.GetDeviceByFingerprintParams{
			UserID:      userID,
			Fingerprint: fingerprint,
		},
	)
}

func (r *Repository) UpdateSession(
	ctx context.Context,
	id pgtype.UUID,
	sessionID pgtype.UUID,
	ip *netip.Addr,
) error {

	return r.q.UpdateDeviceSession(
		ctx,
		db.UpdateDeviceSessionParams{
			ID:        id,
			SessionID: sessionID,
			LastIp:    ip,
		},
	)
}
