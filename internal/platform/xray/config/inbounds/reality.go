package inbounds

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/model"
)

func NewVLESSReality(
	port int,
	privateKey string,
	shortID string,
	serverName string,
	client Client,
) model.Inbound {

	return model.Inbound{

		Tag: "vless-reality",

		Listen: "0.0.0.0",

		Port: port,

		Protocol: "vless",

		Settings: Settings{
			Clients: []Client{
				client,
			},
		},

		StreamSettings: StreamSettings{

			Network: "tcp",

			Security: "reality",

			RealitySettings: &RealitySettings{

				Show: false,

				Dest: serverName + ":443",

				ServerNames: []string{
					serverName,
				},

				PrivateKey: privateKey,

				ShortIds: []string{
					shortID,
				},
			},
		},
	}
}
