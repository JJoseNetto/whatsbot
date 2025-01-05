package middleware

import (
	"net/http"
	"os"
)

func Auth(next http.Handler) http.Handler {
	var apiToken = os.Getenv("API_TOKEN")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		if token != apiToken {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
