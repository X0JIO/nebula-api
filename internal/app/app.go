package app

import (
	"context"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/platform/config"
	"github.com/X0JIO/nebula-api/internal/platform/logger"
	"github.com/X0JIO/nebula-api/internal/platform/web"
	"github.com/X0JIO/nebula-api/internal/platform/database/postgres"

	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	Postgres *postgres.DB
	Server *web.Server
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

	database, err := postgres.New(
	context.Background(),
	cfg.App.Postgres,
	)
	if err != nil {
		return nil, err
	}

	server := web.New(
		cfg.App.Host,
		cfg.App.Port,
	)

	return &App{
		Config: cfg,
		Logger: log,
		Postgres: database,
		Server: server,
	}, nil
	
}

func (a *App) Run(ctx context.Context) error {

	a.Logger.Info("HTTP server starting")

	defer a.Postgres.Close()

	errCh := make(chan error, 1)

	go func() {

		err := a.Server.Start()

		if err != nil && err != http.ErrServerClosed {
			errCh <- err
		}

	}()

	select {

	case <-ctx.Done():

		a.Logger.Info("shutdown signal received")

		return a.Server.Shutdown(context.Background())

	case err := <-errCh:

		return err

	}

}