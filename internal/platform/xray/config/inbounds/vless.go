package inbounds

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/model"
)

func VLESSInbound(port int) model.Inbound {
	return NewVLESSInbound(
		port,
		nil,
	)
}

func NewVLESSInbound(
	port int,
	client *Client,
) model.Inbound {

	clients := []Client{}

	if client != nil {
		clients = append(
			clients,
			*client,
		)
	}

	return BaseInbound(
		"vless",
		"vless",
		port,
		Settings{
			Clients: clients,
		},
		nil,
	)
}
