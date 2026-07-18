package web

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
}

func New(host string, port int) *Server {

	addr := fmt.Sprintf("%s:%d", host, port)

	return &Server{
		http: &http.Server{
			Addr:              addr,
			Handler:           NewRouter(),
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