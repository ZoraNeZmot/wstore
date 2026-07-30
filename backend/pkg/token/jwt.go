package token

import (
	"errors"
	"net/http"
	"strings"
	"time"
	"wstore/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const ClaimsKey = "claims"

type TokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Name   string    `json:"name,omitempty"`
	jwt.RegisteredClaims
}

func NewTokenClaims(userID uuid.UUID, email string, name string, exp time.Time) TokenClaims {
	return TokenClaims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
}

func ValidateAuthHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("unauthorized")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	err := ValidateToken(token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GenerateToken(claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	config := config.GetConfig()
	return token.SignedString([]byte(config.ServerConfig.JwtSecret))
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := config.GetConfig().ServerConfig.JwtSecret
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("Unauthorized")
	}
	return nil
}

func GetClaims(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := config.GetConfig().ServerConfig.JwtSecret
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*TokenClaims), nil
}
