package userhandlers

import (
	"net/http"
	"wstore/internal/config"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/handlers/help/response"
	"wstore/internal/storage"
	"wstore/internal/storage/dbusers"
	tokenpkg "wstore/pkg/token"

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
		"/users",
		handlershelp.AuthMiddleware(u.logger, http.HandlerFunc(u.GetUsersList)),
	)

	mux.Handle("/user/{id}", handlershelp.AuthMiddleware(u.logger, http.HandlerFunc(u.GetUser)))

	return nil
}

// @Summary		Get users list
// @Description	Get users list
// @Tags			user
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.UsersListDTO
// @Failure		401	{object}	response.ErrorDTO
// @Failure		500	{object}	response.ErrorDTO
// @Router			/users [get]
// @Security		Bearer
func (u *UserHandlers) GetUsersList(w http.ResponseWriter, r *http.Request) {
	users, err := u.db.GetUsers(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get users", err.Error())
		u.logger.WithError(err).Error("Failed to get users")
		return
	}
	response.Success(w, http.StatusOK, users)
}

// @Summary		Get user
// @Description	Get user by id or me
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User UUID or me for the current user"	pattern(^(me|[0-9a-fA-F-]{36})$)
// @Success		200	{object}	response.UserDTO
// @Failure		401	{object}	response.ErrorDTO
// @Failure		500	{object}	response.ErrorDTO
// @Router			/user/{id} [get]
// @Security		Bearer
func (u *UserHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var (
		user *dbusers.User
		err  error
	)
	claims, ok := r.Context().Value("claims").(*tokenpkg.TokenClaims)
	if !ok {
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get token claims")
		u.logger.Debug("Failed to get token claims")
		return
	}

	switch id {
	case "me":
		user, err = u.db.GetUserByEmail(r.Context(), dbusers.NewUserWithEmail(claims.Email))
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to get user", err.Error())
			u.logger.WithError(err).Debug("Failed to get user")
			return
		}
		user.Password = ""
	default:
		user, err = u.db.GetUserByUid(r.Context(), dbusers.NewUserWithUid(id))
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to get user", err.Error())
			u.logger.WithError(err).Debug("Failed to get user")
			return
		}
		user.Password = ""
	}
	response.Success(w, http.StatusOK, user)
}
