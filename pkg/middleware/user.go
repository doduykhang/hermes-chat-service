package middleware

import (
	"context"
	"net/http"
)

func GetUserId(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		userIDHeader := r.Header["X-User-Id"]	
		if len(userIDHeader) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("No user id"))
			return
		}

		userID := userIDHeader[0]

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
