package httpserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/X0JIO/nebula-api/internal/modules/admin"
	"github.com/X0JIO/nebula-api/internal/modules/auth"
	"github.com/X0JIO/nebula-api/internal/modules/comments"
	"github.com/X0JIO/nebula-api/internal/modules/devices"
	"github.com/X0JIO/nebula-api/internal/modules/health"
	"github.com/X0JIO/nebula-api/internal/modules/projects"
	"github.com/X0JIO/nebula-api/internal/modules/sessions"
	"github.com/X0JIO/nebula-api/internal/modules/tasks"
	"github.com/X0JIO/nebula-api/internal/modules/users"
	"github.com/X0JIO/nebula-api/internal/modules/vpn"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
	"github.com/X0JIO/nebula-api/internal/platform/xray"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(
	userHandler *users.Handler,
	authHandler *auth.Handler,
	adminHandler *admin.Handler,
	projectsHandler *projects.Handler,
	tasksHandler *tasks.Handler,
	commentsHandler *comments.Handler,
	sessionsHandler *sessions.Handler,
	devicesHandler *devices.Handler,
	vpnHandler *vpn.Handler,
	xrayHandler *xray.HTTPController,
	jwtMiddleware *middleware.JWTMiddleware,
) http.Handler {

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

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

		// protected auth routes
		r.Group(func(r chi.Router) {
			r.Use(jwtMiddleware.Handler)

			r.Post("/auth/logout", authHandler.Logout)
			r.Post("/auth/logout-all", authHandler.LogoutAll)

			r.Get("/users/me", userHandler.Me)
		})

		// authenticated routes

		r.Group(func(r chi.Router) {

			r.Use(
				jwtMiddleware.Handler,
			)
		})

		r.Group(func(r chi.Router) {

			r.Use(jwtMiddleware.Handler)

			r.Route("/projects", func(r chi.Router) {

				r.Get("/", projectsHandler.ListProjects)
				r.Post("/", projectsHandler.CreateProject)

				r.Get("/{id}", projectsHandler.GetProject)
				r.Put("/{id}", projectsHandler.UpdateProject)
				r.Delete("/{id}", projectsHandler.DeleteProject)

				r.Post("/{id}/members", projectsHandler.AddMember)
				r.Delete("/{id}/members/{userId}", projectsHandler.RemoveMember)

			})

		})

		r.Route("/tasks", func(r chi.Router) {

			r.Use(jwtMiddleware.Handler)

			// create task
			r.Post(
				"/",
				tasksHandler.CreateTask,
			)

			// get single task
			r.Get(
				"/{id}",
				tasksHandler.GetTask,
			)

			// update task
			r.Put(
				"/{id}",
				tasksHandler.UpdateTask,
			)

			// delete task
			r.Delete(
				"/{id}",
				tasksHandler.DeleteTask,
			)

			// tasks by project
			r.Get(
				"/project/{projectId}",
				tasksHandler.ListProjectTasks,
			)

			// tasks by assignee
			r.Get(
				"/assignee/{userId}",
				tasksHandler.ListAssigneeTasks,
			)

			// tasks by project and status
			r.Get(
				"/project/{projectId}/status/{status}",
				tasksHandler.ListStatusTasks,
			)

			comments.Routes(
				r,
				commentsHandler,
			)

		})

		r.Route("/devices", func(r chi.Router) {

			r.Use(jwtMiddleware.Handler)

			r.Get("/", devicesHandler.List)

			r.Delete("/", devicesHandler.DeleteAll)

			r.Delete("/{id}", devicesHandler.Delete)
		})

		r.Route("/vpn", func(r chi.Router) {

			r.Use(jwtMiddleware.Handler)

			r.Post("/config", vpnHandler.CreateConfig)

			r.Get("/configs", vpnHandler.ListConfigs)

			r.Get("/subscription", vpnHandler.Subscription)

			r.Get("/subscription/base64", vpnHandler.SubscriptionBase64)
		})

		r.Route("/xray", func(r chi.Router) {

			r.Use(jwtMiddleware.Handler)

			r.Get(
				"/status",
				xrayHandler.Status,
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
