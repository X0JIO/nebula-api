package outbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

func Freedom() model.Outbound {
	return model.Outbound{
		Protocol: "freedom",
		Tag:      "direct",
	}
}
