package inbounds

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/model"
)

func NewVLESSInbound(
	port int,
	client Client,
) model.Inbound {

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

func VLESSInbound(port int) model.Inbound {

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
