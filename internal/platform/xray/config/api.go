package config

type APIConfig struct {
	Tag      string   `json:"tag,omitempty"`
	Services []string `json:"services,omitempty"`
}
