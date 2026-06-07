package handlers

import "net/http"

type Handlers interface {
	RegisterRoutes(mux *http.ServeMux) error
}
