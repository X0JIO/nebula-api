package inbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/config"

type ShadowsocksClient struct {
	Method   string `json:"method"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ShadowsocksSettings struct {
	Clients []ShadowsocksClient `json:"clients"`
}

func NewShadowsocksInbound(
	port int,
	password string,
	email string,
) config.Inbound {

	return BaseInbound(
		"shadowsocks",
		"shadowsocks",
		port,
		ShadowsocksSettings{
			Clients: []ShadowsocksClient{
				{
					Method:   "chacha20-ietf-poly1305",
					Password: password,
					Email:    email,
				},
			},
		},
		nil,
	)
}
