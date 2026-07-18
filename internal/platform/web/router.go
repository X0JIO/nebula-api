package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/X0JIO/nebula-api/internal/modules/auth"
	"github.com/X0JIO/nebula-api/internal/modules/health"
	"github.com/X0JIO/nebula-api/internal/modules/users"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
)

func NewRouter(
	userHandler *users.Handler,
	authHandler *auth.Handler,
	jwtMiddleware *middleware.JWTMiddleware,
) http.Handler {

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Get("/health", health.Handler)

		r.Post("/users", userHandler.Create)

		r.Post("/auth/register", authHandler.Register)

		r.Post("/auth/login", authHandler.Login)

		r.Post("/auth/refresh", authHandler.Refresh)

		r.Group(func(r chi.Router) {

			r.Use(jwtMiddleware.Handler)

			r.Get(
				"/users/me",
				userHandler.Me,
			)

		})

	})
	return r
}
