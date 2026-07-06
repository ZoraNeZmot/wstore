package token

import (
	"errors"
	"net/http"
	"strings"
	"time"
	"wstore/internal/config"
	"wstore/internal/storage"

	"github.com/golang-jwt/jwt/v5"
)

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

func GenerateToken(user *storage.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
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
