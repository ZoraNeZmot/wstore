package authhandlers

import (
	"net/http"
	handlershelp "wstore/internal/handlers/help"
)

func (a *AuthHandlers) EmailRegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc(handlershelp.GetFullPathAPI("auth/email/signup"), a.EmailSignup)
	return nil

}

func (a *AuthHandlers) EmailSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// TODO: Implement
	} else {
		http.Error(w, "EmailSignup: Method not allowed", http.StatusMethodNotAllowed)
		a.logger.Error("EmailSignup: Method not allowed")
		return
	}
	// TODO: Implement
}

func (a *AuthHandlers) EmailLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func (a *AuthHandlers) EmailLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}
