package servicehandlers

import (
	"net/http"
	"strings"
	handlershelp "wstore/internal/handlers/help"
	tokenpkg "wstore/pkg/token"

	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type ServiceHandlers struct {
}

func NewServiceHandlers() *ServiceHandlers {
	return &ServiceHandlers{}
}

func (s *ServiceHandlers) SwaggerLicense(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("Handler Type", "Service").Trace("SwaggerLicense")
	http.ServeFile(w, r, "static/license")
}

// @Summary		Health check
// @Description	Check if the server is running
// @Tags			service
// @Success		200	{string}	string	"OK"
// @Router			/health [get]
func (s *ServiceHandlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("Handler Type", "Service").Trace("HealthCheck")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(handlershelp.JsonOk))
	if err != nil {
		logrus.WithError(err).Error("Failed to write response")
	}
}

// @Summary		Auth test
// @Description	Test if the user is authorized
// @Tags			service
// @Success		200	{string}	string	"OK"
// @Router			/service/authtest [get]
// @Security		Bearer
func (s *ServiceHandlers) AuthTest(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("Handler Type", "Service").Trace("AuthTest")
	token := r.Header.Get("Authorization")
	if token == "" {
		handlershelp.JSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")
	valid, err := tokenpkg.ValidateToken(token)
	if err != nil {
		handlershelp.JSONError(w, http.StatusUnauthorized, "Unauthorized")
		logrus.WithError(err).Warn("AuthTest: Failed to validate token")
		return
	}
	if !valid {
		handlershelp.JSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(handlershelp.JsonOk))
	if err != nil {
		logrus.WithError(err).Error("Failed to write response")
	}

}

func (s *ServiceHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc(handlershelp.GetFullPathAPI("service/health"), s.HealthCheck)
	mux.HandleFunc("/swagger/license.txt", s.SwaggerLicense)
	mux.HandleFunc(handlershelp.GetFullPathAPI("service/authtest"), s.AuthTest)

	// swagger format exception
	mux.Handle("/swagger/", httpSwagger.Handler())

	return nil
}
