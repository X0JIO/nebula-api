package app

import (
	"context"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/modules/auth"
	"github.com/X0JIO/nebula-api/internal/modules/users"
	"github.com/X0JIO/nebula-api/internal/platform/cache/redis"
	"github.com/X0JIO/nebula-api/internal/platform/config"
	"github.com/X0JIO/nebula-api/internal/platform/database/postgres"
	"github.com/X0JIO/nebula-api/internal/platform/logger"
	"github.com/X0JIO/nebula-api/internal/platform/web"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"go.uber.org/zap"
)

type App struct {
	Config      *config.Config
	Logger      *zap.Logger
	Postgres    *postgres.DB
	Redis       *redis.Client
	Users       *users.Service
	Auth        *auth.Service
	UserHandler *users.Handler
	Server      *web.Server
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

	cache, err := redis.New(
		context.Background(),
		cfg.App.Redis,
	)

	queries := db.New(database.Pool)

	userRepository := users.NewRepository(
		queries,
	)

	userService := users.NewService(
		userRepository,
	)

	authService := auth.NewService(
		userRepository,
	)

	userHandler := users.NewHandler(
		userService,
	)

	authHandler := auth.NewHandler(
		authService,
	)

	server := web.New(
		cfg.App.Host,
		cfg.App.Port,
		userHandler,
		authHandler,
	)

	if err != nil {
		return nil, err
	}

	return &App{
		Config:      cfg,
		Logger:      log,
		Postgres:    database,
		Redis:       cache,
		Users:       userService,
		Auth:        authService,
		UserHandler: userHandler,
		Server:      server,
	}, nil

}

func (a *App) Run(ctx context.Context) error {

	a.Logger.Info("HTTP server starting")

	defer a.Postgres.Close()
	defer a.Redis.Close()

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
