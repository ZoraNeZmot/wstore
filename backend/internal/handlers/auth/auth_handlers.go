package authhandlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"
	"wstore/internal/config"
	"wstore/internal/handlers/help/response"
	"wstore/internal/storage"
	"wstore/internal/storage/dbtokens"
	"wstore/pkg/token"

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
	mux.HandleFunc("GET /auth/refresh", a.RefreshToken)
	mux.HandleFunc("POST /auth/logout", a.Logout)
	return nil
}

// @Summary		Refresh token
// @Description	Refresh token. Can't be executed in Swagger UI.
// @Tags			auth
// @Param			Cookie	header	string	true	"Refresh token (format: RefreshToken=YOUR_TOKEN)"
// @Produce		json
// @Success		200	{object}	response.TokenDTO
// @Failure		401	{object}	response.ErrorDTO
// @Failure		500	{object}	response.ErrorDTO
// @Router			/auth/refresh [get]
func (a *AuthHandlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("RefreshToken")
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "Unauthorized", "No cookie RefreshToken present")
		a.logger.Debug("No cookie RefreshToken present")
		return
	}
	if refreshToken.Value == "" {
		response.Error(w, http.StatusUnauthorized, "Unauthorized", "No cookie RefreshToken present")
		a.logger.Debug("Refresh token is empty")
		return
	}
	tokenData := dbtokens.NewTokenWithToken(refreshToken.Value)
	valid, err := a.db.IsTokenValid(r.Context(), tokenData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusUnauthorized, "Unauthorized", "No cookie RefreshToken present")
			a.logger.Debug("Refresh token is not valid or expired")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Server error", err.Error())
		a.logger.WithError(err).Debug("Failed to check token validity")
		return
	}
	if !valid {
		response.Error(w, http.StatusUnauthorized, "Unauthorized", "Invalid refresh token")
		a.logger.Debug("Invalid refresh token")
		return
	}
	claims, err := token.GetClaims(refreshToken.Value)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", err.Error())
		a.logger.WithError(err).Debug("Failed to get token claims")
		return
	}
	accessToken, err := token.GenerateToken(token.NewTokenClaims(claims.UserID, claims.Email, claims.Name, time.Now().Add(time.Minute*15)))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", err.Error())
		a.logger.WithError(err).Debug("Failed to generate access token")
		return
	}
	response.Success(w, http.StatusOK, response.TokenResponse{AccessToken: accessToken})
}

// @Summary		Logout
// @Description	Logout the user.
// @Tags			auth
// @Produce		json
// @Success		200	{object}	response.FlagDTO
// @Failure		401	{object}	response.ErrorDTO
// @Failure		500	{object}	response.ErrorDTO
// @Router			/auth/logout [post]
func (a *AuthHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	clearRefreshCookie := func() {
		http.SetCookie(w, &http.Cookie{
			Name:     "RefreshToken",
			Value:    "",
			Path:     "/api/auth",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})
	}

	refreshToken, err := r.Cookie("RefreshToken")
	if err != nil {
		response.Ok(w)
		a.logger.Debug("No cookie RefreshToken present")
		return
	}
	if refreshToken.Value == "" {
		clearRefreshCookie()
		response.Ok(w)
		a.logger.Debug("Logged out because refresh token is empty")
		return
	}

	tokenData := dbtokens.NewTokenWithToken(refreshToken.Value)
	err = a.db.RevokeToken(r.Context(), tokenData)
	if err != nil {
		a.logger.WithError(err).Debug("Failed to revoke token from database")
	}
	clearRefreshCookie()
	response.Ok(w)
}
