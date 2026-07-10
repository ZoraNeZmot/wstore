package handlershelp

import (
	"net/http"
	"wstore/internal/handlers/help/response"
	tokenpkg "wstore/pkg/token"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := tokenpkg.ValidateAuthHeader(r)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Unauthorized", "Failed to validate token")
			return
		}
		next.ServeHTTP(w, r)
	})
}
