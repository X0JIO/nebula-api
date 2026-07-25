package config

type Inbound struct {
	Tag      string `json:"tag"`
	Listen   string `json:"listen,omitempty"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`

	Settings any `json:"settings,omitempty"`

	StreamSettings any `json:"streamSettings,omitempty"`

	Sniffing any `json:"sniffing,omitempty"`
}
