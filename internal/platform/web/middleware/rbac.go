package middleware

import "net/http"

func RequireRoles(roles ...string) func(http.Handler) http.Handler {

	allowed := make(map[string]struct{}, len(roles))

	for _, role := range roles {
		allowed[role] = struct{}{}
	}

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			role, ok := r.Context().Value(ContextRole).(string)

			if !ok {
				http.Error(
					w,
					"role missing",
					http.StatusForbidden,
				)
				return
			}

			if _, ok := allowed[role]; !ok {

				http.Error(
					w,
					"forbidden",
					http.StatusForbidden,
				)

				return
			}

			next.ServeHTTP(w, r)
		})

	}
}
