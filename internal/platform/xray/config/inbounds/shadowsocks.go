package inbounds

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

type ShadowsocksClient struct {
	Method   string `json:"method"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ShadowsocksSettings struct {
	Clients []ShadowsocksClient `json:"clients"`
}

func Shadowsocks(port int) model.Inbound {

	return BaseInbound(
		"shadowsocks",
		"shadowsocks",
		port,
		ShadowsocksSettings{
			Clients: []ShadowsocksClient{},
		},
		nil,
	)
}
