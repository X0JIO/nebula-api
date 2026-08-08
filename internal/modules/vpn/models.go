package vpn

import "time"

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

type VPN struct {
	ID     string
	UserID string

	Protocol string

	UUID string

	Port int

	PublicKey  string
	PrivateKey string
	ShortID    string

	ServerName string

	Status string

	CreatedAt time.Time
}

type VPNAccountResponse struct {
	ID                string    `json:"id"`
	UUID              string    `json:"uuid"`
	SubscriptionToken string    `json:"subscription_token"`
	PublicKey         string    `json:"public_key"`
	ShortID           string    `json:"short_id"`
	CreatedAt         time.Time `json:"created_at"`
}

type DeleteVPNResponse struct {
	Status string `json:"status"`
}
