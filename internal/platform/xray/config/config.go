package config

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

type Config struct {
	Log     LogConfig     `json:"log,omitempty"`
	API     *APIConfig    `json:"api,omitempty"`
	Stats   *StatsConfig  `json:"stats,omitempty"`
	DNS     DNSConfig     `json:"dns,omitempty"`
	Routing RoutingConfig `json:"routing,omitempty"`
	Policy  PolicyConfig  `json:"policy,omitempty"`

	Inbounds []model.Inbound `json:"inbounds"`

	Outbounds []model.Outbound `json:"outbounds"`
}
