package inbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

type TrojanSettings struct {
	Clients []Client `json:"clients"`
}

func Trojan(port int) config.Inbound {

	return BaseInbound(
		"trojan",
		"trojan",
		port,
		TrojanSettings{
			Clients: []Client{},
		},
		nil,
	)
}
