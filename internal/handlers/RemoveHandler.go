package handlers

import (
	"encoding/json"
	"log"
	"my-api/internal/models"
	"my-api/internal/services"
	"net/http"
)

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	res := models.RemoveResponse{
		Message: "Suppression réussie",
		Status:  200,
	}

	token := r.Header.Get("Authorization")
	_, err := services.Remove(token)

	if err != nil {
		res = models.RemoveResponse{
			Message: "Erreur lors de la suppression",
			Status:  401,
		}
	}

	_, err = services.Disconnect(token)

	if err != nil {
		res = models.RemoveResponse{
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
