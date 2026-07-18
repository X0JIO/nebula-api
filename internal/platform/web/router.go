package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/X0JIO/nebula-api/internal/modules/health"
	"github.com/X0JIO/nebula-api/internal/modules/users"
)

func NewRouter(
	userHandler *users.Handler,
) http.Handler {

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Get("/health", health.Handler)

		r.Post("/users", userHandler.Create)

	})

	return r
}
