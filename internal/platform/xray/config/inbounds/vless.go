package inbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

func NewVLESSInbound(
	port int,
	client Client,
) config.Inbound {

	return BaseInbound(
		"vless",
		"vless",
		port,
		Settings{
			Clients: []Client{
				client,
			},
		},
		nil,
	)
}

func VLESSInbound(port int) config.Inbound {

	return BaseInbound(
		"vless",
		"vless",
		port,
		Settings{
			Clients: []Client{},
		},
		nil,
	)
}
