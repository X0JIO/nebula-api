package auth

import (
	"errors"
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

func (j *JWT) ParseToken(
	tokenString string,
) (string, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			return j.secret, nil
		},
	)

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("invalid claims")
	}

	tokenType, ok := claims["type"].(string)

	if !ok || tokenType != "refresh" {
		return "", errors.New("not refresh token")
	}

	userID, ok := claims["sub"].(string)

	if !ok {
		return "", errors.New("invalid subject")
	}

	return userID, nil
}
