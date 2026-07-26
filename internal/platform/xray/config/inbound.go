package config

type Sniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride,omitempty"`
}
type Inbound struct {
	Tag string `json:"tag"`

	Listen string `json:"listen,omitempty"`

	Port int `json:"port"`

	Protocol string `json:"protocol"`

	Settings any `json:"settings,omitempty"`

	StreamSettings any `json:"streamSettings,omitempty"`

	Sniffing any `json:"sniffing,omitempty"`
}
