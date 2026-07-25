package outbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

func DNS() config.Outbound {
	return config.Outbound{
		Protocol: "dns",
		Tag:      "dns-out",
	}
}
