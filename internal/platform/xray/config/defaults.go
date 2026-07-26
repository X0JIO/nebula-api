package config

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

type Defaults struct {
	Inbounds  []model.Inbound
	Outbounds []model.Outbound
}

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

			Levels: map[string]PolicyLevel{

				"0": {

					StatsUserUplink: true,

					StatsUserDownlink: true,

					StatsUserOnline: true,
				},
			},
		},

		API: &APIConfig{

			Tag: "api",

			Services: []string{

				"HandlerService",

				"StatsService",
			},
		},

		Stats: &StatsConfig{},

		Inbounds: []model.Inbound{},

		Outbounds: []model.Outbound{},
	}
}
