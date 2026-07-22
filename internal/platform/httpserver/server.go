package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/X0JIO/nebula-api/internal/modules/admin"
	"github.com/X0JIO/nebula-api/internal/modules/auth"
	"github.com/X0JIO/nebula-api/internal/modules/comments"
	"github.com/X0JIO/nebula-api/internal/modules/devices"
	"github.com/X0JIO/nebula-api/internal/modules/projects"
	"github.com/X0JIO/nebula-api/internal/modules/sessions"
	"github.com/X0JIO/nebula-api/internal/modules/tasks"
	"github.com/X0JIO/nebula-api/internal/modules/users"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
)

type Server struct {
	http *http.Server
}

func New(
	host string,
	port string,
	userHandler *users.Handler,
	authHandler *auth.Handler,
	adminHandler *admin.Handler,
	projectsHandler *projects.Handler,
	tasksHandler *tasks.Handler,
	commentsHandler *comments.Handler,
	sessionsHandler *sessions.Handler,
	devicesHandler *devices.Handler,
	jwtMiddleware *middleware.JWTMiddleware,
) *Server {

	addr := fmt.Sprintf("%s:%s", host, port)

	return &Server{
		http: &http.Server{
			Addr: addr,
			Handler: NewRouter(
				userHandler,
				authHandler,
				adminHandler,
				projectsHandler,
				tasksHandler,
				commentsHandler,
				sessionsHandler,
				devicesHandler,
				jwtMiddleware,
			),
			ReadHeaderTimeout: 5 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
