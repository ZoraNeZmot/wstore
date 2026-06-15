package authhandlers

import (
	"net/http"
	"wstore/internal/config"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/storage"

	"github.com/sirupsen/logrus"
)

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
	mux.HandleFunc(handlershelp.GetFullPathAPI("auth/google/login"), a.GoogleLogin)
	mux.HandleFunc(handlershelp.GetFullPathAPI("auth/google/callback"), a.GoogleCallback)

	mux.HandleFunc(handlershelp.GetFullPathAPI("auth/email/signup"), a.EmailSignup)
	return nil
}
