package inbounds

type ShadowsocksClient struct {
	Method   string `json:"method"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ShadowsocksSettings struct {
	Clients []ShadowsocksClient `json:"clients"`
}

func Shadowsocks(port int) Inbound {

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
