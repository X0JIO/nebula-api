package inbounds

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/model"
	"github.com/X0JIO/nebula-api/internal/platform/xray/reality"
)

func NewVLESSReality(
	port int,
	credentials *reality.Credentials,
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

				Xver: 0,

				ServerNames: []string{
					serverName,
				},

				PrivateKey: credentials.PrivateKey,

				ShortIds: []string{
					credentials.ShortID,
				},
			},
		},
	}
}
