package vpn

import (
	"context"

	"github.com/google/uuid"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/X0JIO/nebula-api/internal/shared/apperrors"
)

type CreateDeviceRequest struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

func (s *Service) CreateDevice(
	ctx context.Context,
	userID string,
	req CreateDeviceRequest,
) (*db.VpnDevice, error) {

	vpnUser, err := s.repo.GetVPNUser(
		ctx,
		userID,
	)

	if err != nil {
		return nil, err
	}

	device, err := s.repo.CreateVPNDevice(
		ctx,
		db.CreateVPNDeviceParams{
			VpnUserID:   vpnUser.ID,
			Name:        req.Name,
			Platform:    req.Platform,
			DeviceToken: uuid.New().String(),
		},
	)

	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s *Service) ListDevices(
	ctx context.Context,
	userID string,
) ([]db.VpnDevice, error) {

	vpnUser, err := s.repo.GetVPNUser(
		ctx,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return s.repo.ListVPNDevices(
		ctx,
		vpnUser.ID,
	)
}

func (s *Service) DeleteDevice(
	ctx context.Context,
	userID string,
	deviceID string,
) error {

	vpnUser, err := s.repo.GetVPNUser(
		ctx,
		userID,
	)

	if err != nil {
		return err
	}

	id, err := parseUUID(deviceID)

	if err != nil {
		return err
	}

	devices, err := s.repo.ListVPNDevices(
		ctx,
		vpnUser.ID,
	)

	if err != nil {
		return err
	}

	for _, device := range devices {
		if device.ID == id {
			return s.repo.DeleteVPNDevice(
				ctx,
				id,
			)
		}
	}

	return apperrors.ErrVPNUserNotFound
}

func (s *Service) RevokeDevice(
	ctx context.Context,
	userID string,
	deviceID string,
) error {

	_, err := s.repo.GetVPNUser(
		ctx,
		userID,
	)

	if err != nil {
		return err
	}

	id, err := parseUUID(deviceID)

	if err != nil {
		return err
	}

	return s.repo.RevokeVPNDevice(
		ctx,
		id,
	)
}
