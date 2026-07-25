package config

func DefaultOutbounds() []Outbound {
	return []Outbound{
		{
			Tag:      "direct",
			Protocol: "freedom",
		},
		{
			Tag:      "blocked",
			Protocol: "blackhole",
		},
	}
}
