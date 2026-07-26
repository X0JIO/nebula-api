package xray

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/config/inbounds"
)

func (b *Builder) AddDefaultInbounds() {

	b.AddInbound(inbounds.VLESSInbound(443))

	b.AddInbound(inbounds.RealityInbound(8443))

	b.AddInbound(inbounds.VMessInbound(10086))

	b.AddInbound(inbounds.TrojanInbound(4433))

	b.AddInbound(inbounds.ShadowsocksInbound(8388))
}
