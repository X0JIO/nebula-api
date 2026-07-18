package web

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/X0JIO/nebula-api/internal/modules/auth"
	"github.com/X0JIO/nebula-api/internal/modules/users"
)

type Server struct {
	http *http.Server
}

func New(
	host string,
	port int,
	userHandler *users.Handler,
	authHandler *auth.Handler,
) *Server {

	addr := fmt.Sprintf("%s:%d", host, port)

	return &Server{
		http: &http.Server{
			Addr: addr,
			Handler: NewRouter(
				userHandler,
				authHandler,
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
