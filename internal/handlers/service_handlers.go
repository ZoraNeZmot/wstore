package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type ServiceHandlers struct {
}

func NewServiceHandlers() *ServiceHandlers {
	return &ServiceHandlers{}
}

func (s *ServiceHandlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("Handler Type", "Service").Trace("HealthCheck")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		logrus.WithError(err).Error("Failed to write response")
	}
}

func (s *ServiceHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("/health", s.HealthCheck)
	return nil
}
