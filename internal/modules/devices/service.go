package devices

import (
	"context"
	"errors"
	"log"
	"net/netip"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repository *Repository
	sessions   SessionRepository
}

type SessionRepository interface {
	Delete(
		ctx context.Context,
		id pgtype.UUID,
	) error
}

func NewService(
	repository *Repository,
	sessions SessionRepository,
) *Service {

	return &Service{
		repository: repository,
		sessions:   sessions,
	}
}

func (s *Service) List(
	ctx context.Context,
	userID string,
) ([]db.Device, error) {
	log.Printf("devices.List userID=%q", userID)
	var uid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return nil, err
	}

	return s.repository.List(
		ctx,
		uid,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	deviceID string,
) error {

	var id pgtype.UUID

	if err := id.Scan(deviceID); err != nil {
		return err
	}

	device, err := s.repository.Get(ctx, id)
	if err != nil {
		return err
	}

	if device.SessionID.Valid {
		_ = s.sessions.Delete(
			ctx,
			device.SessionID,
		)
	}

	return s.repository.Delete(
		ctx,
		id,
	)
}

func (s *Service) DeleteAll(
	ctx context.Context,
	userID string,
) error {

	var uid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return err
	}

	return s.repository.DeleteUserDevices(
		ctx,
		uid,
	)
}

func (s *Service) Register(
	ctx context.Context,
	userID pgtype.UUID,
	sessionID pgtype.UUID,
	deviceName string,
	platform string,
	fingerprint string,
	vpnUUID pgtype.UUID,
	ip string,
) error {

	log.Printf(
		"Register: user=%s fingerprint=%q",
		userID.String(),
		fingerprint,
	)

	addr, err := netip.ParseAddr(ip)
	if err != nil {
		addr = netip.IPv4Unspecified()
	}

	device, err := s.repository.GetByFingerprint(
		ctx,
		userID,
		fingerprint,
	)

	if err != nil {
		log.Printf("GetByFingerprint error: %v", err)
	} else {
		log.Printf("Found device: %s", device.ID.String())
	}

	if err == nil {
		return s.repository.UpdateSession(
			ctx,
			device.ID,
			sessionID,
			&addr,
		)
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	_, err = s.repository.Create(
		ctx,
		db.CreateDeviceParams{
			UserID:      userID,
			SessionID:   sessionID,
			Name:        deviceName,
			Platform:    platform,
			Fingerprint: fingerprint,
			VpnUuid:     vpnUUID,
			LastIp:      &addr,
		},
	)

	return err
}
