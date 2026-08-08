package vpn

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{
		queries: q,
	}
}

func parseUUID(id string) (pgtype.UUID, error) {
	var pg pgtype.UUID

	if err := pg.Scan(id); err != nil {
		return pgtype.UUID{}, err
	}

	return pg, nil
}

func (r *Repository) GetVPNUser(
	ctx context.Context,
	userID string,
) (db.VpnUser, error) {

	pgID, err := parseUUID(userID)
	if err != nil {
		return db.VpnUser{}, err
	}

	return r.queries.GetVPNUserByUserID(
		ctx,
		pgID,
	)
}

func (r *Repository) CreateVPNUser(
	ctx context.Context,
	userID string,
	userUUID string,
	privateKey string,
	publicKey string,
	shortID string,
	subscriptionToken string,
) (db.VpnUser, error) {

	pgID, err := parseUUID(userID)
	if err != nil {
		return db.VpnUser{}, err
	}

	params := db.CreateVPNUserParams{
		UserID:            pgID,
		Uuid:              userUUID,
		PrivateKey:        privateKey,
		PublicKey:         publicKey,
		ShortID:           shortID,
		SubscriptionToken: subscriptionToken,
	}

	return r.queries.CreateVPNUser(
		ctx,
		params,
	)
}

func (r *Repository) SaveVPNConfig(
	ctx context.Context,
	params db.SaveVPNConfigParams,
) (db.VpnConfig, error) {

	row, err := r.queries.SaveVPNConfig(
		ctx,
		params,
	)

	if err != nil {
		return db.VpnConfig{}, err
	}

	return db.VpnConfig{
		ID:        row.ID,
		VpnUserID: row.VpnUserID,
		DeviceID:  row.DeviceID,
		Protocol:  row.Protocol,
		Config:    row.Config,
		CreatedAt: row.CreatedAt,
	}, nil
}

func (r *Repository) ListVPNConfigs(
	ctx context.Context,
	vpnUserID pgtype.UUID,
) ([]db.VpnConfig, error) {

	return r.queries.ListVPNConfigs(
		ctx,
		vpnUserID,
	)
}

func (r *Repository) GetVPNConfig(
	ctx context.Context,
	id pgtype.UUID,
) (db.VpnConfig, error) {

	return r.queries.GetVPNConfig(
		ctx,
		id,
	)
}

func (r *Repository) DeleteVPNConfig(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.queries.DeleteVPNConfig(
		ctx,
		id,
	)
}

func (r *Repository) GetVPNUserBySubscription(
	ctx context.Context,
	token string,
) (db.VpnUser, error) {

	return r.queries.GetVPNUserBySubscription(
		ctx,
		token,
	)
}

func (r *Repository) DeleteVPNConfigsByUser(
	ctx context.Context,
	vpnUserID pgtype.UUID,
) error {

	return r.queries.DeleteVPNConfigsByUser(
		ctx,
		vpnUserID,
	)
}

func (r *Repository) DeleteVPNUser(
	ctx context.Context,
	userID pgtype.UUID,
) error {

	return r.queries.DeleteVPNUser(
		ctx,
		userID,
	)
}

func (r *Repository) CreateVPNDevice(
	ctx context.Context,
	params db.CreateVPNDeviceParams,
) (db.VpnDevice, error) {

	return r.queries.CreateVPNDevice(
		ctx,
		params,
	)
}

func (r *Repository) ListVPNDevices(
	ctx context.Context,
	vpnUserID pgtype.UUID,
) ([]db.VpnDevice, error) {

	return r.queries.ListVPNDevices(
		ctx,
		vpnUserID,
	)
}

func (r *Repository) RevokeVPNDevice(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.queries.RevokeVPNDevice(
		ctx,
		id,
	)
}

func (r *Repository) DeleteVPNDevice(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.queries.DeleteVPNDevice(
		ctx,
		id,
	)
}

func (r *Repository) GetVPNDevice(
	ctx context.Context,
	id pgtype.UUID,
	vpnUserID pgtype.UUID,
) (db.VpnDevice, error) {

	return r.queries.GetVPNDevice(
		ctx,
		db.GetVPNDeviceParams{
			ID:        id,
			VpnUserID: vpnUserID,
		},
	)
}
