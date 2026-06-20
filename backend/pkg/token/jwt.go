package token

import (
	"errors"
	"time"
	"wstore/internal/config"
	"wstore/internal/storage"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *storage.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	config := config.GetConfig()
	return token.SignedString([]byte(config.ServerConfig.JwtSecret))
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := config.GetConfig().ServerConfig.JwtSecret
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, errors.New("invalid token")
	}
	return true, nil
}
