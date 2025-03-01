package httpHandlers

import (
	"encoding/json"
	"log"
	http3 "my-api/internal/models/http"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	res := http3.HealthResponse{
		Message: "Suppression réussie",
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
