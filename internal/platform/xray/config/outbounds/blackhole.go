package outbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

func Blackhole() model.Outbound {
	return model.Outbound{
		Protocol: "blackhole",
		Tag:      "blocked",
	}
}
