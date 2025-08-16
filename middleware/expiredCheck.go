package middleware

import (
	"config-service/auth"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

func IsExpired(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.Contains(authHeader, "Bearer ") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"message": "Unauthorized",
					"error":   "No token provided",
				})
				return
			}

			token := strings.Split(authHeader, "Bearer ")[1]

			currentUser, err := auth.ValidateToken(token, jwtSecret)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"message": "Unauthorized",
					"error":   err.Error(),
				})
				return
			}

			// Add user to request context
			ctx := context.WithValue(r.Context(), "currentUser", currentUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
