package xray

import (
	"context"

	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
	"github.com/X0JIO/nebula-api/internal/platform/xray/config/inbounds"
)

type ConfigService struct {
	builder *config.Builder
	writer  ConfigWriter
}

func NewConfigService(
	writer ConfigWriter,
) *ConfigService {

	return &ConfigService{

		builder: config.NewBuilder(),

		writer: writer,
	}
}

func (s *ConfigService) Build() (config.Config, error) {

	return s.builder.Build()
}

func (s *ConfigService) Generate(
	ctx context.Context,
) error {

	cfg, err := s.Build()

	if err != nil {
		return err
	}

	data, err := config.Generate(cfg)

	if err != nil {
		return err
	}

	return s.writer.Save(
		ctx,
		data,
	)
}

func (s *ConfigService) AddRealityInbound(
	port int,
	privateKey string,
	shortID string,
	serverName string,
	client inbounds.Client,
) {

	s.builder.AddInbound(

		inbounds.NewVLESSReality(

			port,

			privateKey,

			shortID,

			serverName,

			client,
		),
	)
}
