package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMiddleware struct {
	secret []byte
}

func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{
		secret: []byte(secret),
	}
}

func (m *JWTMiddleware) Handler(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(
				w,
				"missing authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(
				w,
				"invalid authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				return m.secret, nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(
				w,
				"invalid token",
				http.StatusUnauthorized,
			)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(
				w,
				"invalid claims",
				http.StatusUnauthorized,
			)
			return
		}

		userID, ok := claims["sub"].(string)

		if !ok {
			http.Error(
				w,
				"invalid user",
				http.StatusUnauthorized,
			)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			"user_id",
			userID,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
