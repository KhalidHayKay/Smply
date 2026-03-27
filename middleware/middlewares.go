package middleware

import (
	"log"
	"net/http"
	"smply/handler"
	"smply/service"
)

func RequireKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if key == "" {
			handler.Error(w, http.StatusUnauthorized, "API key required")
			return
		}

		valid, err := service.ValidateAPIKey(r.Context(), key)
		if err != nil {
			log.Println("Error validating API key:", err)
			handler.Error(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		if !valid {
			handler.Error(w, http.StatusUnauthorized, "Invalid API key")
			return
		}

		next.ServeHTTP(w, r)
	})
}
