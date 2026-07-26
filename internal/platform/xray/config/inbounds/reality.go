package inbounds

type RealitySettings struct {
	Show        bool     `json:"show"`
	Dest        string   `json:"dest"`
	Xver        int      `json:"xver"`
	ServerNames []string `json:"serverNames"`

	PrivateKey string   `json:"privateKey"`
	ShortIds   []string `json:"shortIds"`
}

type StreamSettings struct {
	Network  string `json:"network"`
	Security string `json:"security,omitempty"`

	RealitySettings *RealitySettings `json:"realitySettings,omitempty"`
}

type Sniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride,omitempty"`
}

type Inbound struct {
	Tag      string `json:"tag"`
	Listen   string `json:"listen,omitempty"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`

	Settings any `json:"settings,omitempty"`

	StreamSettings StreamSettings `json:"streamSettings,omitempty"`

	Sniffing Sniffing `json:"sniffing,omitempty"`
}

type Settings struct {
	Clients []Client `json:"clients,omitempty"`
}

type Client struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
}

func BaseInbound(
	tag string,
	protocol string,
	port int,
	settings Settings,
	stream StreamSettings,
) Inbound {

	return Inbound{
		Tag:      tag,
		Listen:   "0.0.0.0",
		Port:     port,
		Protocol: protocol,

		Settings: settings,

		StreamSettings: stream,
	}
}

func NewVLESSReality(
	port int,
	privateKey string,
	shortID string,
	serverName string,
	client Client,
) Inbound {

	return Inbound{

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
