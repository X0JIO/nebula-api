package xray

import (
	"path/filepath"

	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
)

func NewDefaultConfigService(
	basePath string,
) *ConfigService {

	writer := NewConfigWriter(

		filepath.Join(
			basePath,
			"xray",
			"config.json",
		),
	)

	service := NewConfigService(
		writer,
	)

	service.builder.EnableAPI(
		config.APIConfig{
			Tag: "api",
			Services: []string{
				"HandlerService",
				"StatsService",
			},
		},
	)

	service.builder.EnableStats()

	return service
}
