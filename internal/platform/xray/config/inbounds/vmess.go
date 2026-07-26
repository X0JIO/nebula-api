package inbounds

import (
	"github.com/X0JIO/nebula-api/internal/platform/xray/model"
)

type VMessClient struct {
	ID       string `json:"id"`
	AlterID  int    `json:"alterId"`
	Email    string `json:"email"`
	Security string `json:"security"`
}

type VMessSettings struct {
	Clients []VMessClient `json:"clients"`
}

type VMessTCPSettings struct{}

type VMessWSSettings struct {
	Path    string            `json:"path"`
	Headers map[string]string `json:"headers,omitempty"`
}

type VMessRealitySettings struct {
	Show        bool     `json:"show"`
	Dest        string   `json:"dest"`
	Xver        int      `json:"xver"`
	ServerNames []string `json:"serverNames"`
	PrivateKey  string   `json:"privateKey"`
	ShortIds    []string `json:"shortIds"`
}

type VMessStreamSettings struct {
	Network         string                `json:"network"`
	Security        string                `json:"security,omitempty"`
	TCPSettings     *VMessTCPSettings     `json:"tcpSettings,omitempty"`
	WSSettings      *VMessWSSettings      `json:"wsSettings,omitempty"`
	RealitySettings *VMessRealitySettings `json:"realitySettings,omitempty"`
}

type VMessSniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride"`
	RouteOnly    bool     `json:"routeOnly,omitempty"`
}

func NewVMessInbound(
	tag string,
	port int,
	uuid string,
	email string,
) model.Inbound {

	return model.Inbound{
		Tag:      tag,
		Listen:   "0.0.0.0",
		Port:     port,
		Protocol: "vmess",

		Settings: VMessSettings{
			Clients: []VMessClient{
				{
					ID:       uuid,
					AlterID:  0,
					Email:    email,
					Security: "auto",
				},
			},
		},

		StreamSettings: VMessStreamSettings{
			Network: "tcp",
		},

		Sniffing: VMessSniffing{
			Enabled: true,
			DestOverride: []string{
				"http",
				"tls",
				"quic",
			},
		},
	}
}

func VMessInbound(port int) model.Inbound {

	return NewVMessInbound(
		"vmess",
		port,
		"",
		"",
	)
}
