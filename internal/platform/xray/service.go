package xray

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
	"github.com/X0JIO/nebula-api/internal/platform/xray/config/inbounds"
)

type ConfigService struct {
	builder *config.Builder
}

func NewConfigService() *ConfigService {
	return &ConfigService{
		builder: config.NewBuilder(),
	}
}

func (s *ConfigService) BuildVLESSReality(
	port int,
	privateKey string,
	shortID string,
	serverName string,
	client inbounds.Client,
) ([]byte, error) {

	s.builder.AddInbound(
		inbounds.NewVLESSReality(
			port,
			privateKey,
			shortID,
			serverName,
			client,
		),
	)

	cfg := s.builder.Config()

	return config.Generate(cfg)
}
