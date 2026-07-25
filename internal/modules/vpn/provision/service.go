package provision

import (
	"context"

	"github.com/X0JIO/nebula-api/internal/platform/xray/reality"
)

type Service struct {
	ports PortAllocator
}

func NewService(
	ports PortAllocator,
) *Service {

	return &Service{
		ports: ports,
	}
}

func (s *Service) Create(
	ctx context.Context,
) (*ProvisionResult, error) {

	port, err := s.ports.Allocate(ctx)
	if err != nil {
		return nil, err
	}

	cfg, err := reality.Generate()
	if err != nil {
		return nil, err
	}

	return &ProvisionResult{
		UUID:       cfg.UUID,
		Port:       port,
		PrivateKey: cfg.PrivateKey,
		PublicKey:  cfg.PublicKey,
		ShortID:    cfg.ShortID,
	}, nil
}
