package config

type Outbound struct {
	Tag      string `json:"tag"`
	Protocol string `json:"protocol"`

	Settings any `json:"settings,omitempty"`
}
