package config

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/config/inbounds"
	"github.com/X0JIO/nebula-api/internal/platform/xray/config/outbounds"
)

func DefaultVPNConfig() Config {

	builder := NewBuilder()

	builder.SetLog("warning")

	builder.AddOutbound(
		outbounds.Freedom(),
		outbounds.Blackhole(),
	)

	builder.AddInbound(
		inbounds.VLESSInbound(443),
	)

	builder.AddInbound(
		inbounds.VMessInbound(10086),
	)

	builder.AddInbound(
		inbounds.TrojanInbound(4433),
	)

	builder.AddInbound(
		inbounds.ShadowsocksInbound(8388),
	)

	builder.AddInbound(
		inbounds.RealityInbound(8443),
	)

	cfg := builder.Config()

	return cfg
}
