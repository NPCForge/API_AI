package handlers

import (
	"encoding/json"
	"log"
	"my-api/internal/models"
	"net/http"
)

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	res := models.DisconnectResponse{
		Message: "Deconnexion réussie",
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
