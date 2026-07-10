package handlers

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Use outer mux for swagger handlers(no http.StripPrefix)
func RegisterSwagger(mux *http.ServeMux) error {
	mux.HandleFunc("/swagger/license.txt", SwaggerLicenseHandler)
	mux.Handle("/swagger/", httpSwagger.Handler())
	return nil
}

func SwaggerLicenseHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/license")
}
