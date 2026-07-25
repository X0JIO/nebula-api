package config

func Default() Config {
	return Config{
		Log: LogConfig{
			Loglevel: "warning",
		},

		DNS: DNSConfig{
			Servers: []string{
				"1.1.1.1",
				"8.8.8.8",
			},
		},

		Routing: RoutingConfig{
			DomainStrategy: "AsIs",
			Rules:          []RoutingRule{},
		},

		Policy: PolicyConfig{
			Levels: map[string]PolicyLevel{},
		},

		API: nil,

		Stats: nil,

		Inbounds: []Inbound{},

		Outbounds: []Outbound{
			{
				Tag:      "direct",
				Protocol: "freedom",
			},
			{
				Tag:      "blocked",
				Protocol: "blackhole",
			},
		},
	}
}
