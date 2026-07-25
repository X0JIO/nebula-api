package xray

func NewClient(cfg *Config) Client {

	if cfg == nil {
		return NewStubClient()
	}

	if !cfg.Enabled {
		return NewStubClient()
	}

	// Пока используем HTTP API.
	return NewHTTPClient(cfg)
}
