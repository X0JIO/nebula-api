package middleware

import (
	"context"
	"net/http"
)

type ProjectRoleChecker interface {
	GetProjectRole(
		ctx context.Context,
		userID string,
		projectID string,
	) (string, error)
}

type ProjectRoleMiddleware struct {
	checker ProjectRoleChecker
}

func NewProjectRoleMiddleware(
	checker ProjectRoleChecker,
) *ProjectRoleMiddleware {

	return &ProjectRoleMiddleware{
		checker: checker,
	}
}

func (m *ProjectRoleMiddleware) Require(
	roles ...string,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			userID := UserID(r.Context())
			projectID := ProjectID(r.Context())

			if userID == "" || projectID == "" {

				http.Error(
					w,
					"unauthorized",
					http.StatusUnauthorized,
				)

				return
			}

			role, err := m.checker.GetProjectRole(
				r.Context(),
				userID,
				projectID,
			)

			if err != nil {

				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)

				return
			}

			for _, allowed := range roles {

				if role == allowed {

					next.ServeHTTP(
						w,
						r,
					)

					return
				}
			}

			http.Error(
				w,
				"forbidden",
				http.StatusForbidden,
			)
		})
	}
}
