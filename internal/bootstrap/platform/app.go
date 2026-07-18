package bootstrap

import (
	"context"

	"github.com/your-org/nebula-api/internal/platform/config"
	"github.com/your-org/nebula-api/internal/platform/logger"
)

type App struct {
	Config *config.Config
	Logger logger.Logger
}

func New(ctx context.Context) (*App, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	log, err := logger.New(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		Logger: log,
	}, nil
}

func (a *App) Run(ctx context.Context) error {

	a.Logger.Info("Nebula API started")

	select {}

}