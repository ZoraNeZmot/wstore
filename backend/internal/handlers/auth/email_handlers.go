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

	"github.com/google/uuid"
)

type EmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type EmailResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

var OkResponse = EmailResponse{
	Message: "ok",
	Error:   "",
}

func JSONError(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EmailResponse{
		Error: err,
	})
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
// @Success		201		{object}	EmailResponse
// @Failure		400		{object}	EmailResponse
// @Failure		500		{object}	EmailResponse
// @Router			/auth/email/signup [post]
func (a *AuthHandlers) EmailSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		a.logger.Error("EmailSignup: Method not allowed")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		JSONError(w, http.StatusBadRequest, "Failed to read body")
		a.logger.Errorf("EmailSignup: Failed to read body: %v", err)
		return
	}
	defer r.Body.Close()
	var user EmailRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		JSONError(w, http.StatusBadRequest, "Failed to parse json body")
		a.logger.Errorf("EmailSignup: Failed to unmarshal body: %v", err)
		return
	}

	argon := argon2.New(nil, nil)
	hashPassword, err := argon.Hash(user.Password)
	if err != nil {
		JSONError(w, http.StatusInternalServerError, "Failed to hash password")
		a.logger.Errorf("EmailSignup: Failed to hash password: %v", err)
		return
	}

	user.Password = string(hashPassword)
	user.Email = strings.ToLower(user.Email)
	user.Name = strings.ToUpper(user.Name)
	id := uuid.New()
	userData := storage.NewUser(id, user.Email, user.Password, user.Name)
	exists := a.db.CheckUser(r.Context(), userData)
	if exists {
		JSONError(w, http.StatusBadRequest, dberrors.UserAlreadyExists.Error())
		a.logger.Warnf("EmailSignup: User already exists: %v", err)
		return
	}

	_, err = a.db.AddUser(
		r.Context(),
		userData,
	)
	if err != nil {
		JSONError(w, http.StatusInternalServerError, "Failed to add user. Try again later")
		a.logger.Errorf("EmailSignup: Failed to add user: %v", err)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(OkResponse)
	}
}

func (a *AuthHandlers) EmailLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		a.logger.Error("EmailLogin: Method not allowed")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		JSONError(w, http.StatusBadRequest, "Failed to read body")
		a.logger.Errorf("EmailLogin: Failed to read body: %v", err)
		return
	}
	defer r.Body.Close()
	var user EmailRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		JSONError(w, http.StatusBadRequest, "Failed to parse json body")
		a.logger.Errorf("EmailLogin: Failed to unmarshal body: %v", err)
		return
	}

	user.Email = strings.ToLower(user.Email)

	userData, err := a.db.GetUserByEmail(r.Context(), storage.NewUserWithEmail(user.Email))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			JSONError(w, http.StatusUnauthorized, "User not found. Check your email or password")
			a.logger.Errorf("EmailLogin: User not found: %v", err)
			return
		}

		JSONError(w, http.StatusInternalServerError, "Failed to get user. Try again later")
		a.logger.Errorf("EmailLogin: Failed to get user: %v", err)
		return
	}

	isValid, err := argon2.Verify(user.Password, userData.Password)
	if err != nil {
		JSONError(w, http.StatusInternalServerError, "Failed to verify password. Try again later")
		a.logger.Errorf("EmailLogin: Failed to verify password: %v", err)
		return
	}
	if !isValid {
		JSONError(w, http.StatusUnauthorized, "Invalid email or password")
		a.logger.Warnf("EmailLogin: Invalid email or password") // TODO: later change log level to info
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(OkResponse)
}

func (a *AuthHandlers) EmailLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement should I do that here?
}
