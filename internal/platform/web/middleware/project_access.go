package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProjectAccessChecker interface {
	IsProjectMember(
		ctx context.Context,
		userID string,
		projectID string,
	) (bool, error)
}

type ProjectAccessMiddleware struct {
	checker ProjectAccessChecker
}

func NewProjectAccessMiddleware(
	checker ProjectAccessChecker,
) *ProjectAccessMiddleware {

	return &ProjectAccessMiddleware{
		checker: checker,
	}
}

func (m *ProjectAccessMiddleware) Handler(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		projectID := chi.URLParam(
			r,
			"projectID",
		)

		if projectID == "" {

			http.Error(
				w,
				"missing project id",
				http.StatusBadRequest,
			)

			return
		}

		userID := UserID(
			r.Context(),
		)

		if userID == "" {

			http.Error(
				w,
				"unauthorized",
				http.StatusUnauthorized,
			)

			return
		}

		ok, err := m.checker.IsProjectMember(
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

		if !ok {

			http.Error(
				w,
				"forbidden",
				http.StatusForbidden,
			)

			return
		}

		ctx := WithProjectID(
			r.Context(),
			projectID,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
