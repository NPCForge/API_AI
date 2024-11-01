package services

import (
	"my-api/pkg"
	"net/http"
)

// LoggingMiddleware est un middleware qui journalise chaque requête HTTP.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupère le token de l'en-tête Authorization
		token := r.Header.Get("Authorization")

		// Vérifie si le token est présent
		if token == "" {
			http.Error(w, "Token manquant dans l'en-tête Authorization", http.StatusUnauthorized)
			return
		}

		// Vérifie si le token est valide (cette étape dépend de ton application)
		if !pkg.IsValidToken(token) {
			http.Error(w, "Token non valide", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r) // Appel du prochain handler
	})
}
