package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/X0JIO/nebula-api/internal/modules/admin"
	"github.com/X0JIO/nebula-api/internal/modules/auth"
	"github.com/X0JIO/nebula-api/internal/modules/health"
	"github.com/X0JIO/nebula-api/internal/modules/users"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
)

func NewRouter(
	userHandler *users.Handler,
	authHandler *auth.Handler,
	adminHandler *admin.Handler,
	jwtMiddleware *middleware.JWTMiddleware,
) http.Handler {

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		// public routes

		r.Get(
			"/health",
			health.Handler,
		)

		r.Post(
			"/auth/register",
			authHandler.Register,
		)

		r.Post(
			"/auth/login",
			authHandler.Login,
		)

		r.Post(
			"/auth/refresh",
			authHandler.Refresh,
		)

		// authenticated routes

		r.Group(func(r chi.Router) {

			r.Use(
				jwtMiddleware.Handler,
			)

			r.Get(
				"/users/me",
				userHandler.Me,
			)

		})

		// admin routes

		r.Group(func(r chi.Router) {

			r.Use(
				jwtMiddleware.Handler,
			)

			r.Use(
				middleware.RequireRoles("admin"),
			)

			r.Get(
				"/admin/dashboard",
				adminHandler.Dashboard,
			)

			r.Get(
				"/admin/users",
				adminHandler.ListUsers,
			)

			r.Get(
				"/admin/users/{id}",
				adminHandler.GetUser,
			)

			r.Patch(
				"/admin/users/{id}/role",
				adminHandler.ChangeRole,
			)

			r.Patch(
				"/admin/users/{id}/status",
				adminHandler.ChangeStatus,
			)

			r.Delete(
				"/admin/users/{id}",
				adminHandler.DeleteUser,
			)

		})

	})

	return r
}
