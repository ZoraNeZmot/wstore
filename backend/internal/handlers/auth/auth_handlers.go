package authhandlers

import (
	"net/http"
	"wstore/internal/config"
	"wstore/internal/storage"

	"github.com/sirupsen/logrus"
)

type EmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AuthHandlers struct {
	db     *storage.Database
	Config *config.Config
	logger *logrus.Entry
}

func NewAuthHandlers() *AuthHandlers {
	config := config.GetConfig()
	return &AuthHandlers{
		db:     storage.NewDatabase(config.DatabaseConfig.GetDB()),
		Config: config,
		logger: logrus.WithField("Handler Type", "Auth"),
	}
}

func (a *AuthHandlers) RegisterRoutes(mux *http.ServeMux) error {
	err := a.EmailRegisterRoutes(mux)
	if err != nil {
		return err
	}
	// err = a.GoogleRegisterRoutes(mux)
	// if err != nil {
	// 	return err
	// }

	return nil
}
