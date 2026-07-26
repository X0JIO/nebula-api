package config

import "time"

type XRayConfig struct {
	Enabled bool

	BaseURL string
	APIKey  string

	Timeout time.Duration

	BinaryPath string
	ConfigPath string
	WorkingDir string
}
