package config

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

func DefaultOutbounds() []model.Outbound {
	return []model.Outbound{
		{
			Tag:      "direct",
			Protocol: "freedom",
		},
		{
			Tag:      "blocked",
			Protocol: "blackhole",
		},
	}
}
