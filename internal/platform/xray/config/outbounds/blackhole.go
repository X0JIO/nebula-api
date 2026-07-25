package outbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

func Blackhole() config.Outbound {
	return config.Outbound{
		Protocol: "blackhole",
		Tag:      "blocked",
	}
}
