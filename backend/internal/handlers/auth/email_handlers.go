package authhandlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/storage"
	"wstore/pkg/argon2"

	"github.com/google/uuid"
)

type EmailSignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (a *AuthHandlers) EmailRegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc(handlershelp.GetFullPathAPI("auth/email/signup"), a.EmailSignup)
	return nil

}

func (a *AuthHandlers) EmailSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "EmailSignup: Failed to read body", http.StatusBadRequest)
			a.logger.Errorf("EmailSignup: Failed to read body: %v", err)
			return
		}
		defer r.Body.Close()
		var user EmailSignupRequest
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "EmailSignup: Failed to unmarshal body", http.StatusBadRequest)
			a.logger.Errorf("EmailSignup: Failed to unmarshal body: %v", err)
			return
		}
		argon := argon2.New(nil, nil)
		hashPassword, err := argon.Hash(user.Password)
		if err != nil {
			http.Error(w, "EmailSignup: Failed to hash password", http.StatusInternalServerError)
			a.logger.Errorf("EmailSignup: Failed to hash password: %v", err)
			return
		}
		user.Password = string(hashPassword)
		user.Email = strings.ToLower(user.Email)
		user.Name = strings.ToUpper(user.Name)
		id := uuid.New().String()
		_, err = a.db.AddUser(
			r.Context(),
			storage.NewUser(id, user.Email, user.Password, user.Name),
		)
		if err != nil {
			http.Error(w, "EmailSignup: Failed to add user", http.StatusInternalServerError)
			a.logger.Errorf("EmailSignup: Failed to add user: %v", err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok")) //TODO: email verification
	} else {
		http.Error(w, "EmailSignup: Method not allowed", http.StatusMethodNotAllowed)
		a.logger.Error("EmailSignup: Method not allowed")
		return
	}

}

func (a *AuthHandlers) EmailLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func (a *AuthHandlers) EmailLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}
