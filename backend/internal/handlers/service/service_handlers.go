package servicehandlers

import (
	"net/http"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/handlers/help/response"
)

type ServiceHandlers struct {
}

func NewServiceHandlers() *ServiceHandlers {
	return &ServiceHandlers{}
}

// Register routes for the service handlers
func (s *ServiceHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("/service/health", s.HealthCheck)
	mux.Handle(
		"/service/authtest",
		handlershelp.AuthMiddleware(http.HandlerFunc(s.AuthTest)),
	)

	return nil
}

// @Summary		Health check
// @Description	Check if the server is running
// @Tags			service
// @Success		200	{object}	response.FlagDTO
// @Example		{success: true}
// @Router			/service/health [get]
func (s *ServiceHandlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.Ok(w)
}

// @Summary		Auth test
// @Description	Test if the user is authorized
// @Tags			service
// @Success		200	{object}	response.FlagDTO
// @Example		{success: true}
// @Router			/service/authtest [get]
// @Security		Bearer
func (s *ServiceHandlers) AuthTest(w http.ResponseWriter, r *http.Request) {
	response.Ok(w)
}
