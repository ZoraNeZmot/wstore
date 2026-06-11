package authhandlers

import (
	"net/http"
	"wstore/internal/config"
	handlershelp "wstore/internal/handlers/help"

	"github.com/sirupsen/logrus"
)

type AuthHandlers struct {
	Config *config.Config
	logger *logrus.Entry
}

func NewAuthHandlers() *AuthHandlers {
	return &AuthHandlers{
		Config: config.GetConfig(),
		logger: logrus.WithField("Handler Type", "Auth"),
	}
}

func (a *AuthHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc(handlershelp.GetFullPathAPI("auth/google/login"), a.GoogleLogin)
	mux.HandleFunc(handlershelp.GetFullPathAPI("auth/google/callback"), a.GoogleCallback)
	return nil
}

// func (a *AuthHandlers) GoogleLogin(w http.ResponseWriter, r *http.Request) {
// 	state := uuid.New().String()
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "state",
// 		Value:    state,
// 		Path:     "/",
// 		MaxAge:   0,
// 		HttpOnly: true,
// 		SameSite: http.SameSiteLaxMode,
// 	})
// 	url := a.Config.GoogleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)

// 	a.logger.WithField("Handler Name", "GoogleLogin").Debugf("Google state: %s", state)
// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
// }

// func (a *AuthHandlers) GoogleCallback(w http.ResponseWriter, r *http.Request) {
// 	state := r.URL.Query().Get("state")
// 	code := r.URL.Query().Get("code")
// 	a.logger.WithField("Handler Name", "GoogleCallback").Debugf("GoogleCallback state: %s, code: %s", state, code)

// 	if state == "" {
// 		http.Error(w, "State not found", http.StatusBadRequest)
// 		a.logger.WithField("Handler Name", "GoogleCallback").Error("State not found")
// 		return
// 	}
// 	cookie, err := r.Cookie("state")
// 	if err != nil {
// 		http.Error(w, "State not found", http.StatusBadRequest)
// 		a.logger.WithField("Handler Name", "GoogleCallback").Error("State not found in COOKIE")
// 		return
// 	}
// 	if cookie.Value != state {
// 		http.Error(w, "State mismatch", http.StatusBadRequest)
// 		a.logger.WithField("Handler Name", "GoogleCallback").Error("State mismatch")
// 		return
// 	}
// 	_, err = a.Config.GoogleOAuthConfig.Exchange(r.Context(), code) // TODO: TOKEN store
// 	if err != nil {
// 		http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
// 		a.logger.WithField("Handler Name", "GoogleCallback").Error("Failed to exchange code")
// 		return
// 	}

// 	logrus.WithField("Handler Type", "Auth").WithField("Handler Name", "GoogleCallback").Debug("SUCCESS exchange code for token")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Success"))
// }
