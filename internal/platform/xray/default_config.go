package xray

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/config"
	"github.com/X0JIO/nebula-api/internal/platform/xray/config/inbounds"
)

func DefaultConfig() config.Config {

	builder := config.NewBuilder()

	builder.AddInbound(
		inbounds.VLESSInbound(443),
	)

	builder.AddInbound(
		inbounds.NewVLESSReality(
			8443,
			"",
			"",
			"example.com",
			inbounds.Client{},
		),
	)

	builder.AddInbound(
		inbounds.VMessInbound(10086),
	)

	builder.AddInbound(
		inbounds.Trojan(4433),
	)

	builder.AddInbound(
		inbounds.Shadowsocks(8388),
	)

	return builder.Config()
}
