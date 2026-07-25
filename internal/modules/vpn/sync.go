package vpn

import (
	"context"
	"fmt"
	"sync"

	"github.com/X0JIO/nebula-api/internal/platform/xray"
)

type SyncService struct {
	client xray.Client

	mu       sync.RWMutex
	inbounds map[string]int
}

func NewSyncService(client xray.Client) *SyncService {
	return &SyncService{
		client:   client,
		inbounds: make(map[string]int),
	}
}

func (s *SyncService) inboundID(
	ctx context.Context,
	protocol string,
) (int, error) {

	s.mu.RLock()
	id, ok := s.inbounds[protocol]
	s.mu.RUnlock()

	if ok {
		return id, nil
	}

	inbounds, err := s.client.ListInbounds(ctx)
	if err != nil {
		return 0, err
	}

	for _, inbound := range inbounds {

		s.mu.Lock()
		s.inbounds[inbound.Protocol] = inbound.ID
		s.mu.Unlock()

		if inbound.Protocol == protocol {
			return inbound.ID, nil
		}
	}

	return 0, fmt.Errorf("xray inbound %q not found", protocol)
}

func (s *SyncService) EnsureInbound(
	ctx context.Context,
	protocol string,
) error {

	_, err := s.inboundID(ctx, protocol)

	return err
}

func (s *SyncService) AddUser(
	ctx context.Context,
	uuid string,
	email string,
	protocol string,
) error {

	if _, err := s.inboundID(ctx, protocol); err != nil {
		return err
	}

	return s.client.AddUser(
		ctx,
		xray.AddUserRequest{
			UUID:     uuid,
			Email:    email,
			Protocol: protocol,
			Inbound:  protocol,
		},
	)
}

func (s *SyncService) UpdateUser(
	ctx context.Context,
	uuid string,
	enabled bool,
	expiresAt int64,
) error {

	return s.client.UpdateUser(
		ctx,
		xray.UpdateUserRequest{
			UUID:      uuid,
			Enabled:   enabled,
			ExpiresAt: expiresAt,
		},
	)
}

func (s *SyncService) DeleteUser(
	ctx context.Context,
	uuid string,
) error {

	return s.client.DeleteUser(
		ctx,
		uuid,
	)
}

func (s *SyncService) Stats(
	ctx context.Context,
	uuid string,
) (*xray.UserStats, error) {

	return s.client.GetUserStats(
		ctx,
		uuid,
	)
}

func (s *SyncService) Refresh(
	ctx context.Context,
) error {

	inbounds, err := s.client.ListInbounds(ctx)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.inbounds = make(map[string]int)

	for _, inbound := range inbounds {
		s.inbounds[inbound.Protocol] = inbound.ID
	}

	return nil
}
