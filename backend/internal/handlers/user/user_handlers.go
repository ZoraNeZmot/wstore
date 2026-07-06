package userhandlers

import (
	"encoding/json"
	"net/http"
	"wstore/internal/config"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/storage"

	"github.com/sirupsen/logrus"
)

type UserHandlers struct {
	db     *storage.Database
	Config *config.Config
	logger *logrus.Entry
}

func NewUserHandlers() *UserHandlers {
	config := config.GetConfig()
	return &UserHandlers{
		db:     storage.NewDatabase(config.DatabaseConfig.GetDB()),
		Config: config,
		logger: logrus.WithField("Handler Type", "User"),
	}
}

func (u *UserHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.Handle(
		handlershelp.GetFullPathAPI("users"),
		handlershelp.AuthMiddleware(http.HandlerFunc(u.GetUsers)),
	)
	return nil
}

func (u *UserHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.db.GetUsers(r.Context())
	if err != nil {
		handlershelp.JSONError(w, http.StatusInternalServerError, "Failed to get users")
		u.logger.WithError(err).Error("Failed to get users")
		return
	}
	json.NewEncoder(w).Encode(users)
}
