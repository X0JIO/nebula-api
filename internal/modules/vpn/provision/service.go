package provision

import (
	"context"

	"github.com/X0JIO/nebula-api/internal/platform/xray"
	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
)

type Service struct {
	xray      xray.Client
	generator *config.Generator
	writer    xray.ConfigWriter
}

func NewService(
	client xray.Client,
	generator *config.Generator,
	writer xray.ConfigWriter,
) *Service {
	return &Service{
		xray:      client,
		generator: generator,
		writer:    writer,
	}
}

func (s *Service) Add(
	ctx context.Context,
	protocol string,
	uuid string,
	email string,
) error {

	switch protocol {

	case "vless":
		return s.addVLESS(ctx, uuid, email)

	case "reality":
		return s.addReality(ctx, uuid, email)

	case "vmess":
		return s.addVMess(ctx, uuid, email)

	case "trojan":
		return s.addTrojan(ctx, uuid, email)

	case "shadowsocks":
		return s.addShadowsocks(ctx, uuid, email)

	default:
		return xray.ErrUnsupportedProtocol
	}
}
