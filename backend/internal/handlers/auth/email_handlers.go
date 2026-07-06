package authhandlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/storage"
	"wstore/internal/storage/dberrors"
	"wstore/pkg/argon2"
	"wstore/pkg/token"
)

type EmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (a *AuthHandlers) EmailRegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc(handlershelp.GetFullPathAPI("auth/email/signup"), a.EmailSignup)
	return nil

}

// @Summary		Email signup
// @Description	Signup with email and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			user	body		EmailRequest	true	"User"
// @Success		201		{object}	EmailMessage
// @Failure		400		{object}	EmailError
// @Failure		500		{object}	EmailError
// @Router			/auth/email/signup [post]
func (a *AuthHandlers) EmailSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handlershelp.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		a.logger.Error("EmailSignup: Method not allowed")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handlershelp.JSONError(w, http.StatusBadRequest, "Failed to read body")
		a.logger.Errorf("EmailSignup: Failed to read body: %v", err)
		return
	}
	defer r.Body.Close()
	var user EmailRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		handlershelp.JSONError(w, http.StatusBadRequest, "Failed to parse json body")
		a.logger.Errorf("EmailSignup: Failed to unmarshal body: %v", err)
		return
	}

	argon := argon2.New(nil, nil)
	hashPassword, err := argon.Hash(user.Password)
	if err != nil {
		handlershelp.JSONError(w, http.StatusInternalServerError, "Failed to hash password")
		a.logger.Errorf("EmailSignup: Failed to hash password: %v", err)
		return
	}

	user.Password = string(hashPassword)
	user.Email = strings.ToLower(user.Email)
	user.Name = strings.ToUpper(user.Name)

	userData := storage.NewUser(user.Email, user.Password, user.Name)
	exists := a.db.IsUserExists(r.Context(), userData)
	if exists {
		handlershelp.JSONError(w, http.StatusBadRequest, dberrors.UserAlreadyExists.Error())
		a.logger.Warnf("EmailSignup: User already exists: %v", err)
		return
	}

	_, err = a.db.CreateUser(
		r.Context(),
		userData,
	)
	if err != nil {
		handlershelp.JSONError(w, http.StatusInternalServerError, "Failed to add user. Try again later")
		a.logger.Errorf("EmailSignup: Failed to add user: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(handlershelp.JsonOk))
	if err != nil {
		a.logger.Warnf("EmailSignup: Failed to write response: %v", err)
		return
	}
	// TODO: validate email
}

// @Summary		Email login
// @Description	Login with email and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			user	body		EmailRequest	true	"User"
// @Success		200		{object}	EmailMessage
// @Failure		400		{object}	EmailError
// @Failure		500		{object}	EmailError
// @Router			/auth/email/login [post]
func (a *AuthHandlers) EmailLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handlershelp.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		a.logger.Error("EmailLogin: Method not allowed")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		handlershelp.JSONError(w, http.StatusBadRequest, "Failed to read body")
		a.logger.Errorf("EmailLogin: Failed to read body: %v", err)
		return
	}
	defer r.Body.Close()
	var user EmailRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		handlershelp.JSONError(w, http.StatusBadRequest, "Failed to parse json body")
		a.logger.Errorf("EmailLogin: Failed to unmarshal body: %v", err)
		return
	}

	user.Email = strings.ToLower(user.Email)

	userData, err := a.db.GetUserByEmail(r.Context(), storage.NewUserWithEmail(user.Email))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlershelp.JSONError(w, http.StatusUnauthorized, "User not found. Check your email or password")
			a.logger.Errorf("EmailLogin: User not found: %v", err)
			return
		}

		handlershelp.JSONError(w, http.StatusInternalServerError, "Failed to get user. Try again later")
		a.logger.Errorf("EmailLogin: Failed to get user: %v", err)
		return
	}

	isValid, err := argon2.Verify(user.Password, userData.Password)
	if err != nil {
		handlershelp.JSONError(w, http.StatusInternalServerError, "Failed to verify password. Try again later")
		a.logger.Errorf("EmailLogin: Failed to verify password: %v", err)
		return
	}
	if !isValid {
		handlershelp.JSONError(w, http.StatusUnauthorized, "Invalid email or password")
		a.logger.Warnf("EmailLogin: Invalid email or password") // TODO: later change log level to info
		return
	}

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(OkResponse)
	tokenString, err := token.GenerateToken(userData)
	if err != nil {
		handlershelp.JSONError(w, http.StatusInternalServerError, "Failed to generate token. Try again later")
		a.logger.Errorf("EmailLogin: Failed to generate token: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handlershelp.JsonResponseToken{Token: tokenString})
}

func (a *AuthHandlers) EmailLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement should I do that here?
}
