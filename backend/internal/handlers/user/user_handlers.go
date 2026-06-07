package userhandlers

import (
	"net/http"
	handlershelp "wstore/internal/handlers/help"
)

type UserHandlers struct {
	// TODO: Add DB and config
}

func NewUserHandlers() *UserHandlers {
	return &UserHandlers{}
}

func (u *UserHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc(handlershelp.GetFullPathAPI("users"), u.GetUsers)
	return nil
}

func (u *UserHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not implemented"))
}
