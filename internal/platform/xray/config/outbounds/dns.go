package outbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

func DNS() model.Inbound {
	return model.Inbound{
		Protocol: "dns",
		Tag:      "dns-out",
	}
}
