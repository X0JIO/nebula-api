package config

type PolicyConfig struct {
	Levels map[string]PolicyLevel `json:"levels,omitempty"`
}

type PolicyLevel struct {
	Handshake    int `json:"handshake,omitempty"`
	ConnIdle     int `json:"connIdle,omitempty"`
	UplinkOnly   int `json:"uplinkOnly,omitempty"`
	DownlinkOnly int `json:"downlinkOnly,omitempty"`
}
