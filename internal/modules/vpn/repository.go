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

	return r.queries.SaveVPNConfig(
		ctx,
		params,
	)
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
