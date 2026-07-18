package app

import (
	"context"
	"fmt"

	"github.com/X0JIO/nebula-api/internal/platform/config"
)

type App struct {
	cfg *config.Config
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &App{
		cfg: cfg,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	fmt.Printf(
		"%s started on %s:%d\n",
		a.cfg.AppName,
		a.cfg.ServerHost,
		a.cfg.ServerPort,
	)

	<-ctx.Done()

	return nil
}