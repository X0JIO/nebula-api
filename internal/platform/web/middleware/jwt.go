package middleware

import (
	"context"
	"net/http"
	"strings"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Subject   string `json:"sub"`
	Role      string `json:"role"`
	SessionID string `json:"sid"`
	Type      string `json:"type"`
	jwt.RegisteredClaims
}

type SessionRepository interface {
	GetByID(
		context.Context,
		string,
	) (db.Session, error)
}

type JWTMiddleware struct {
	secret   []byte
	sessions SessionRepository
}

func NewJWTMiddleware(
	secret string,
	sessions SessionRepository,
) *JWTMiddleware {

	return &JWTMiddleware{
		secret:   []byte(secret),
		sessions: sessions,
	}
}

func (m *JWTMiddleware) Handler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(
				w,
				"missing authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		parts := strings.SplitN(header, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(
				w,
				"invalid authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(
			parts[1],
			claims,
			func(token *jwt.Token) (interface{}, error) {

				// Разрешаем только HMAC (HS256)
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrTokenSignatureInvalid
				}

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

		if claims.Type != "access" {
			http.Error(
				w,
				"access token required",
				http.StatusUnauthorized,
			)

			return
		}

		session, err := m.sessions.GetByID(
			r.Context(),
			claims.SessionID,
		)
		if err != nil {
			http.Error(
				w,
				"session not found",
				http.StatusUnauthorized,
			)
			return
		}

		if session.Revoked {
			http.Error(
				w,
				"session revoked",
				http.StatusUnauthorized,
			)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			ContextUserID,
			claims.Subject,
		)

		ctx = context.WithValue(
			ctx,
			ContextSessionID,
			claims.SessionID,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})

}
