package config

type Config struct {
	Log       LogConfig     `json:"log,omitempty"`
	API       *APIConfig    `json:"api,omitempty"`
	Stats     *StatsConfig  `json:"stats,omitempty"`
	DNS       DNSConfig     `json:"dns,omitempty"`
	Routing   RoutingConfig `json:"routing,omitempty"`
	Policy    PolicyConfig  `json:"policy,omitempty"`
	Inbounds  []Inbound     `json:"inbounds"`
	Outbounds []Outbound    `json:"outbounds"`
}
