package config

import "encoding/json"

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

func (b *Builder) Build() ([]byte, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}

	return json.MarshalIndent(b.cfg, "", "  ")
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

func (b *Builder) AddInbound(inbound Inbound) {
	b.cfg.Inbounds = append(b.cfg.Inbounds, inbound)
}

func (b *Builder) AddOutbounds(outbounds ...Outbound) {
	b.cfg.Outbounds = append(b.cfg.Outbounds, outbounds...)
}

func (b *Builder) Validate() error {
	for _, inbound := range b.cfg.Inbounds {
		if err := ValidateInbound(inbound); err != nil {
			return err
		}
	}

	return nil
}
