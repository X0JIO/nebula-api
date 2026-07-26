package config

import "github.com/X0JIO/nebula-api/internal/platform/xray/model"

type Builder struct {
	cfg Config
}

func NewBuilder() *Builder {
	return &Builder{
		cfg: Default(),
	}
}

func (b *Builder) Config() Config {
	return b.cfg
}

func (b *Builder) Build() (Config, error) {
	if err := b.Validate(); err != nil {
		return Config{}, err
	}

	return b.cfg, nil
}

func (b *Builder) SetLog(level string) {
	b.cfg.Log.Loglevel = level
}

func (b *Builder) SetDNS(dns DNSConfig) {
	b.cfg.DNS = dns
}

func (b *Builder) SetRouting(r RoutingConfig) {
	b.cfg.Routing = r
}

func (b *Builder) SetPolicy(p PolicyConfig) {
	b.cfg.Policy = p
}

func (b *Builder) EnableAPI(api APIConfig) {
	b.cfg.API = &api
}

func (b *Builder) EnableStats() {
	b.cfg.Stats = &StatsConfig{}
}

func (b *Builder) AddInbound(inbound model.Inbound) {
	b.cfg.Inbounds = append(b.cfg.Inbounds, inbound)
}

func (b *Builder) AddOutbound(outbound model.Outbound) {
	b.cfg.Outbounds = append(
		b.cfg.Outbounds,
		outbound,
	)
}

func (b *Builder) Validate() error {
	for _, inbound := range b.cfg.Inbounds {
		if err := ValidateInbound(inbound); err != nil {
			return err
		}
	}

	return nil
}
