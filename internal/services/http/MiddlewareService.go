package httpServices

import (
	"my-api/pkg"
	"net/http"
)

// LoggingMiddleware is a middleware that logs each HTTP request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the token from the Authorization header
		token := r.Header.Get("Authorization")

		// Check if the token is present
		if token == "" {
			http.Error(w, "Missing token in the Authorization header", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid (this step depends on your application)
		if !pkg.IsValidToken(token) {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r) // Call the next handler
	})
}
