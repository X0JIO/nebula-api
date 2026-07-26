package inbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

type TrojanSettings struct {
	Clients []Client `json:"clients"`
}

func NewTrojanInbound(
	port int,
	client Client,
) config.Inbound {

	return BaseInbound(
		"trojan",
		"trojan",
		port,
		TrojanSettings{
			Clients: []Client{
				client,
			},
		},
		nil,
	)
}
