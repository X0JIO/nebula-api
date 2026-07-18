package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName string

	Environment string

	ServerHost string

	ServerPort string

	LogLevel string
}

func Load() (*Config, error) {

	viper.SetConfigFile(".env")

	_ = viper.ReadInConfig()

	viper.AutomaticEnv()

	cfg := &Config{

		AppName: viper.GetString("APP_NAME"),

		Environment: viper.GetString("APP_ENV"),

		ServerHost: viper.GetString("SERVER_HOST"),

		ServerPort: viper.GetString("SERVER_PORT"),

		LogLevel: viper.GetString("LOG_LEVEL"),
	}

	return cfg, nil

}