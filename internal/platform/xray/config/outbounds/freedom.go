package outbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

func Freedom() config.Outbound {
	return config.Outbound{
		Protocol: "freedom",
		Tag:      "direct",
	}
}
