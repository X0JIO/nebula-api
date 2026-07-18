package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret []byte
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: []byte(secret),
	}
}

func (j *JWT) GenerateAccessToken(
	userID string,
	ttl time.Duration,
) (string, error) {

	claims := jwt.MapClaims{
		"sub":  userID,
		"type": "access",
		"exp":  time.Now().Add(ttl).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(j.secret)
}

func (j *JWT) GenerateRefreshToken(
	userID string,
	ttl time.Duration,
) (string, error) {

	claims := jwt.MapClaims{
		"sub":  userID,
		"type": "refresh",
		"exp":  time.Now().Add(ttl).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(j.secret)
}
