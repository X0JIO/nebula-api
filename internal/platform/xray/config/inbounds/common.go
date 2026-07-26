package inbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

type Settings struct {
	Clients []Client `json:"clients,omitempty"`
}

type Client struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Method   string `json:"method,omitempty"`
}

type StreamSettings struct {
	Network  string `json:"network,omitempty"`
	Security string `json:"security,omitempty"`

	RealitySettings *RealitySettings `json:"realitySettings,omitempty"`
}

type RealitySettings struct {
	Show        bool     `json:"show"`
	Dest        string   `json:"dest"`
	Xver        int      `json:"xver"`
	ServerNames []string `json:"serverNames"`

	PrivateKey string   `json:"privateKey"`
	ShortIds   []string `json:"shortIds"`
}

func BaseInbound(
	tag string,
	protocol string,
	port int,
	settings any,
	stream any,
) model.Inbound {

	return model.Inbound{
		Tag:      tag,
		Listen:   "0.0.0.0",
		Port:     port,
		Protocol: protocol,

		Settings:       settings,
		StreamSettings: stream,
	}
}
