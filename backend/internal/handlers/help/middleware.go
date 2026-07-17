package handlershelp

import (
	"context"
	"net/http"
	"wstore/internal/handlers/help/response"
	tokenpkg "wstore/pkg/token"

	"github.com/sirupsen/logrus"
)

func AuthMiddleware(logger *logrus.Entry, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := tokenpkg.ValidateAuthHeader(r)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Unauthorized", "Invalid token")
			logger.WithError(err).Debug("Cannot validate token")
			return
		}
		claims, err := tokenpkg.GetClaims(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Unauthorized", "Invalid token")
			logger.WithError(err).WithField("token", token).Debug("Cannot get token claims")
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "claims", claims))
		next.ServeHTTP(w, r)
	})
}
