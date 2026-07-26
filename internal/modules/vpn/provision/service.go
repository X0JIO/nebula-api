package provision

import (
	"context"

	"github.com/X0JIO/nebula-api/internal/platform/xray"
)

type Service struct {
	xray xray.Client
}

func NewService(client xray.Client) *Service {
	return &Service{
		xray: client,
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
