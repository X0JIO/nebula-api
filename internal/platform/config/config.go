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

	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	Database string `env:"POSTGRES_DB" envDefault:"nebula"`
	User     string `env:"POSTGRES_USER" envDefault:"nebula"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"nebula"`
	SSLMode  string `env:"POSTGRES_SSLMODE" envDefault:"disable"`

	MaxConns int32 `env:"POSTGRES_MAX_CONNS" envDefault:"20"`
	MinConns int32 `env:"POSTGRES_MIN_CONNS" envDefault:"2"`
}

func Load() (*Config, error) {

	cfg := &Config{}

	if err := env.Parse(&cfg.App); err != nil {
		return nil, err
	}

	return cfg, nil
}