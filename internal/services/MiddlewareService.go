package services

import (
	"encoding/json"
	"my-api/pkg"
	"net/http"

	"github.com/gorilla/websocket"
)

func LoginMiddlewareWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) bool {
	var msg struct {
		Action string `json:"action"`
		Token  string `json:"token"`
	}

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return false
	}

	if msg.Token == "" {
		sendError(conn, "No token in request body")
		return false
	}
	if !pkg.IsValidToken(msg.Token) {
		sendError(conn, "Invalid Token")
		return false
	}
	return true
}

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
