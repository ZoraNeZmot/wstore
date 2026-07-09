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

type TokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Name   string    `json:"name"`
	Exp    time.Time `json:"exp"`
}

func ValidateAuthHeader(r *http.Request) error {
	token := r.Header.Get("Authorization")
	if token == "" {
		return errors.New("unauthorized")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	err := ValidateToken(token)
	if err != nil {
		return err
	}
	return nil
}

func GenerateToken(claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   claims.UserID,
		"email": claims.Email,
		"name":  claims.Name,
		"exp":   claims.Exp.Unix(),
	})
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
