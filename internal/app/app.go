package app

import (
	"github.com/X0JIO/nebula-api/internal/platform/config"
	"github.com/X0JIO/nebula-api/internal/platform/logger"

	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
}

func New() (*App, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	log, err := logger.New(cfg.App.LogLevel)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		Logger: log,
	}, nil

}