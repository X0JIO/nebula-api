package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/X0JIO/nebula-api/internal/modules/health"
)

func NewRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/health", health.Handler)

	return r
}