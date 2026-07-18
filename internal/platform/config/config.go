package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Name string `env:"APP_NAME,required"`
	Env  string `env:"APP_ENV" envDefault:"development"`

	Host string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	Port int    `env:"SERVER_PORT" envDefault:"8080"`

	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func Load() (*Config, error) {

	cfg := &Config{}

	if err := env.Parse(&cfg.App); err != nil {
		return nil, err
	}

	return cfg, nil
}