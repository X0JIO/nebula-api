package config

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/config/inbounds"
	"github.com/X0JIO/nebula-api/internal/platform/xray/config/outbounds"
	"github.com/X0JIO/nebula-api/internal/platform/xray/reality"
)

func DefaultVPNConfig() Config {

	builder := NewBuilder()

	builder.SetLog("warning")

	builder.AddOutbound(
		outbounds.Freedom(),
	)

	builder.AddOutbound(
		outbounds.Blackhole(),
	)

	builder.AddInbound(
		inbounds.VLESSInbound(443),
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

	credentials, err := reality.Generate()

	if err != nil {
		panic(err)
	}

	builder.AddInbound(
		inbounds.NewVLESSReality(
			8443,
			credentials,
			"example.com",
			inbounds.Client{
				ID:    credentials.UUID,
				Email: "default",
			},
		),
	)

	cfg := builder.Config()

	return cfg
}
