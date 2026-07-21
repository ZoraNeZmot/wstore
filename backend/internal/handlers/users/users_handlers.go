package usershandlers

import (
	"net/http"
	"os"
	"wstore/internal/config"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/handlers/help/response"
	"wstore/internal/storage"
	"wstore/internal/storage/dbusers"
	tokenpkg "wstore/pkg/token"

	"github.com/sirupsen/logrus"
)

type UsersHandlers struct {
	db     *storage.Database
	Config *config.Config
	logger *logrus.Entry
}

func NewUsersHandlers() *UsersHandlers {
	config := config.GetConfig()
	return &UsersHandlers{
		db:     storage.NewDatabase(config.DatabaseConfig.GetDB()),
		Config: config,
		logger: logrus.WithField("Handler Type", "Users"),
	}
}

func (u *UsersHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.Handle(
		"/users",
		handlershelp.AuthMiddleware(u.logger, http.HandlerFunc(u.GetUsersList)),
	)

	mux.Handle("/user/{id}", handlershelp.AuthMiddleware(u.logger, http.HandlerFunc(u.GetUser)))
	mux.Handle("/user/profile-image/{id}", handlershelp.AuthMiddleware(u.logger, http.HandlerFunc(u.GetUserProfileImage)))
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
func (u *UsersHandlers) GetUsersList(w http.ResponseWriter, r *http.Request) {
	users, err := u.db.GetUsers(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot get users list")
		u.logger.WithError(err).Error("Cannot get users list")
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
func (u *UsersHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var (
		user *dbusers.User
		err  error
	)

	switch id {
	case "me":
		claims, ok := r.Context().Value("claims").(*tokenpkg.TokenClaims)
		if !ok {
			response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get token claims")
			u.logger.Debug("Failed to get token claims")
			return
		}
		user, err = u.db.GetUserByEmail(r.Context(), dbusers.NewUserWithEmail(claims.Email))
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Server error", "Cannot get user")
			u.logger.WithError(err).Debug("Cannot get user")
			return
		}
		user.Password = ""
	default:
		user, err = u.db.GetUserByid(r.Context(), dbusers.NewUserWithid(id))
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Server error", "Cannot get user")
			u.logger.WithError(err).Debug("Cannot get user")
			return
		}
		user.Password = ""
	}
	response.Success(w, http.StatusOK, user)
}

// @Summary		Get user profile image
// @Description	Get user profile image by id or me
// @Tags			user
// @Accept			json
// @Produce		image/png
// @Param			id	path		string	true	"User UUID or me for the current user"	pattern(^(me|[0-9a-fA-F-]{36})$)
// @Success		200	{binary}	[]byte image/jpeg
// @Failure		401	{object}	response.ErrorDTO
// @Failure		500	{object}	response.ErrorDTO
// @Router			/user/profile-image/{id} [get]
// @Security		Bearer
func (u *UsersHandlers) GetUserProfileImage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	switch id {
	case "me":
		claims, ok := r.Context().Value("claims").(*tokenpkg.TokenClaims)
		if !ok {
			response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get token claims")
			u.logger.Debug("Failed to get token claims")
			return
		}
		imagePath, err := u.db.GetUserProfileImage(r.Context(), dbusers.NewUserWithid(claims.UserID))
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Server error", "Cannot get user profile image")
			u.logger.WithError(err).Debug("Cannot get user profile image. User not found.")
			return
		}
		info, err := os.Stat(imagePath)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Server error", "Cannot get user profile image")
			u.logger.WithError(err).Debug("Cannot get user profile image. File does not exist.")
			return
		}
		if info.IsDir() {
			response.Error(w, http.StatusInternalServerError, "Server error", "Cannot get user profile image")
			u.logger.WithError(err).Debug("Cannot get user profile image. File is a directory.")
			return
		}
		image, err := os.ReadFile(imagePath)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Server error", "Cannot read user profile image")
			u.logger.WithError(err).Debug("Cannot read user profile image")
			return
		}
		response.ImageResponse(w, http.StatusOK, image)
	default:
	}

}
