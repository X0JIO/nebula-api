package inbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

func NewVLESSReality(
	port int,
	privateKey string,
	shortID string,
	serverName string,
	client Client,
) config.Inbound {

	return config.Inbound{

		Tag: "vless",

		Port: port,

		Listen: "0.0.0.0",

		Protocol: "vless",

		Settings: Settings{
			Clients: []Client{
				client,
			},
		},

		StreamSettings: StreamSettings{

			Network: "tcp",

			Security: "reality",

			RealitySettings: RealitySettings{

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

		Sniffing: Sniffing{

			Enabled: true,

			DestOverride: []string{
				"http",
				"tls",
				"quic",
			},
		},
	}
}
