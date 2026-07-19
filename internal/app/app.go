package app

import (
	"context"
	"net/http"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/X0JIO/nebula-api/internal/modules/auth"
	"github.com/X0JIO/nebula-api/internal/modules/users"

	"github.com/X0JIO/nebula-api/internal/modules/admin"
	"github.com/X0JIO/nebula-api/internal/modules/projects"
	"github.com/X0JIO/nebula-api/internal/modules/tasks"
	"github.com/X0JIO/nebula-api/internal/platform/cache/redis"
	"github.com/X0JIO/nebula-api/internal/platform/config"
	"github.com/X0JIO/nebula-api/internal/platform/database/postgres"
	"github.com/X0JIO/nebula-api/internal/platform/httpserver"
	"github.com/X0JIO/nebula-api/internal/platform/logger"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"

	"go.uber.org/zap"
)

type App struct {
	Config   *config.Config
	Logger   *zap.Logger
	Postgres *postgres.DB
	Redis    *redis.Client

	Users *users.Service
	Auth  *auth.Service

	UserHandler     *users.Handler
	AdminHandler    *admin.Handler
	ProjectsHandler *projects.Handler
	TasksHandler    *tasks.Handler
	Server          *httpserver.Server
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
	if err != nil {
		return nil, err
	}

	queries := db.New(database.Pool)

	userRepository := users.NewRepository(queries)

	authRepository := auth.NewRepository(queries)

	userService := users.NewService(userRepository)

	jwt := auth.NewJWT(cfg.App.JWT.Secret)

	authService := auth.NewService(
		userRepository,
		authRepository,
		jwt,
		cfg.App.JWT,
	)

	userHandler := users.NewHandler(userService)

	adminRepository := admin.NewRepository(queries)

	adminService := admin.NewService(adminRepository)

	adminHandler := admin.NewHandler(adminService)

	authHandler := auth.NewHandler(authService)

	projectsRepository := projects.NewRepository(
		queries,
	)

	projectsService := projects.NewService(
		projectsRepository,
	)

	projectsHandler := projects.NewHandler(
		projectsService,
	)

	tasksRepository := tasks.NewRepository(
		queries,
	)

	tasksService := tasks.NewService(
		tasksRepository,
	)

	tasksHandler := tasks.NewHandler(
		tasksService,
	)

	jwtMiddleware := middleware.NewJWTMiddleware(
		cfg.App.JWT.Secret,
	)

	server := httpserver.New(
		cfg.App.Host,
		cfg.App.Port,
		userHandler,
		authHandler,
		adminHandler,
		projectsHandler,
		tasksHandler,
		jwtMiddleware,
	)

	return &App{
		Config:          cfg,
		Logger:          log,
		Postgres:        database,
		Redis:           cache,
		Users:           userService,
		Auth:            authService,
		UserHandler:     userHandler,
		AdminHandler:    adminHandler,
		ProjectsHandler: projectsHandler,
		TasksHandler:    tasksHandler,
		Server:          server,
	}, nil
}

func (a *App) Run(ctx context.Context) error {

	a.Logger.Info("HTTP server starting")

	defer a.Postgres.Close()
	defer a.Redis.Close()

	errCh := make(chan error, 1)

	go func() {
		if err := a.Server.Start(); err != nil && err != http.ErrServerClosed {
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
