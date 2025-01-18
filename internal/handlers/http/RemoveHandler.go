package httpHandlers

import (
	"encoding/json"
	"log"
	http3 "my-api/internal/models/http"
	http2 "my-api/internal/services/http"
	"net/http"
)

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RemoveHandler")
	res := http3.RemoveResponse{
		Message: "Suppression réussie",
		Status:  200,
	}

	token := r.Header.Get("Authorization")
	_, err := http2.Remove(token)

	if err != nil {
		res = http3.RemoveResponse{
			Message: "Erreur lors de la suppression",
			Status:  401,
		}
	}

	_, err = http2.Disconnect(token)

	if err != nil {
		res = http3.RemoveResponse{
			Message: "Erreur lors de la deconnexion",
			Status:  401,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
