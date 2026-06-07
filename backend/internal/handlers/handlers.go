package handlers

<<<<<<< HEAD
import "net/http"
=======
import (
	"net/http"
	authhandlers "wstore/internal/handlers/auth"
	servicehandlers "wstore/internal/handlers/service"
	userhandlers "wstore/internal/handlers/user"
)
>>>>>>> swag

type Handlers interface {
	RegisterRoutes(mux *http.ServeMux) error
}
<<<<<<< HEAD
=======

func RegisterRoutes(mux *http.ServeMux) error {
	authHandlers := authhandlers.NewAuthHandlers()
	if err := authHandlers.RegisterRoutes(mux); err != nil {
		return err
	}
	serviceHandlers := servicehandlers.NewServiceHandlers()
	if err := serviceHandlers.RegisterRoutes(mux); err != nil {
		return err
	}
	userHandlers := userhandlers.NewUserHandlers()
	if err := userHandlers.RegisterRoutes(mux); err != nil {
		return err
	}
	return nil
}
>>>>>>> swag
