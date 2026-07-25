package vpn

type CreateVPNResponse struct {
	UserID            string `json:"user_id"`
	SubscriptionToken string `json:"subscription_token"`
}

type CreateConfigResponse struct {
	ID       string `json:"id"`
	Protocol string `json:"protocol"`
	Config   string `json:"config"`
}

type ConfigResponse struct {
	VLESS       string `json:"vless,omitempty"`
	Reality     string `json:"reality,omitempty"`
	VMess       string `json:"vmess,omitempty"`
	Trojan      string `json:"trojan,omitempty"`
	Shadowsocks string `json:"shadowsocks,omitempty"`
}

type SubscriptionResponse struct {
	URL string `json:"url"`
}

type CreateConfigRequest struct {
	Protocol string `json:"protocol"`
}
