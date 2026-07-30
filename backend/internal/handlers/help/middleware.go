package handlershelp

import (
	"context"
	"net/http"
	"wstore/internal/handlers/help/response"
	"wstore/internal/storage"
	"wstore/internal/storage/dbitems"
	tokenpkg "wstore/pkg/token"

	"github.com/sirupsen/logrus"
)

// TODO: Can think about using middleware as struct methods

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
		r = r.WithContext(context.WithValue(r.Context(), tokenpkg.ClaimsKey, claims))
		next.ServeHTTP(w, r)
	})
}

// func ItemsAuthMiddlewareT(logger *logrus.Entry, db *storage.Database, next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token, err := tokenpkg.ValidateAuthHeader(r)
// 		if err != nil {
// 			response.Error(w, http.StatusUnauthorized, "Unauthorized", "Invalid token")
// 			logger.WithError(err).Debug("Cannot validate token")
// 			return
// 		}
// 	})
// }

func ItemsAuthMiddleware(logger *logrus.Entry, db *storage.Database, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetClaimsFromContext(r)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Unauthorized", "Invalid token")
			logger.WithError(err).Debug("Cannot get token claims")
			return
		}
		itemID, err := GetUUIDFromPath(r, "id")
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get item ID from path")
			logger.WithError(err).Debug("Failed to get item ID from path")
			return
		}
		item := dbitems.NewItem(dbitems.WithID(itemID), dbitems.WithSellerID(claims.UserID))
		isOwner, err := db.CheckItemOwnership(r.Context(), item)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Server error", "Cannot check item ownership")
			logger.WithError(err).Debug("Cannot check item ownership")
			return
		}
		if !isOwner {
			response.Error(w, http.StatusForbidden, "Forbidden", "You are not authorized to access this resource")
			logger.Debug("User is not authorized to access this resource")
			return
		}
		next.ServeHTTP(w, r)
	})
}
