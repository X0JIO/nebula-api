package inbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

type Client struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`

	Flow string `json:"flow,omitempty"`

	Password string `json:"password,omitempty"`

	Method string `json:"method,omitempty"`
}

type Settings struct {
	Clients []Client `json:"clients"`
}

type Sniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride"`
}

func BaseInbound(
	tag string,
	protocol string,
	port int,
	settings any,
	stream any,
) config.Inbound {
	return config.Inbound{
		Tag:            tag,
		Listen:         "0.0.0.0",
		Port:           port,
		Protocol:       protocol,
		Settings:       settings,
		StreamSettings: stream,
		Sniffing: &config.Sniffing{
			Enabled:      true,
			DestOverride: []string{"http", "tls"},
		},
	}
}
