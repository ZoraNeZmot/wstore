package handlers

import (
	"net/http"
	authhandlers "wstore/internal/handlers/auth"
	saleshandlers "wstore/internal/handlers/sales"
	servicehandlers "wstore/internal/handlers/service"
	usershandlers "wstore/internal/handlers/users"
)

type Handlers interface {
	RegisterRoutes(mux *http.ServeMux) error
}

func RegisterRoutes(mux *http.ServeMux) error {
	authHandlers := authhandlers.NewAuthHandlers()
	if err := authHandlers.RegisterRoutes(mux); err != nil {
		return err
	}
	serviceHandlers := servicehandlers.NewServiceHandlers()
	if err := serviceHandlers.RegisterRoutes(mux); err != nil {
		return err
	}
	userHandlers := usershandlers.NewUsersHandlers()
	if err := userHandlers.RegisterRoutes(mux); err != nil {
		return err
	}
	salesHandlers := saleshandlers.NewSalesHandlers()
	if err := salesHandlers.RegisterRoutes(mux); err != nil {
		return err
	}
	return nil
}
