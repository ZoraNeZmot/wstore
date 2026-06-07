package servicehandlers

import (
	"net/http"
	handlershelp "wstore/internal/handlers/help"

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
	_, err := w.Write([]byte("OK"))
	if err != nil {
		logrus.WithError(err).Error("Failed to write response")
	}
}

func (s *ServiceHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc(handlershelp.GetFullPathAPI("health"), s.HealthCheck)
	mux.HandleFunc("/swagger/license.txt", s.SwaggerLicense)

	// swagger format exception
	mux.Handle("/swagger/", httpSwagger.Handler())

	return nil
}
