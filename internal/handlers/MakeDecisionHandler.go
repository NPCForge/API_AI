package handlers

import (
	"encoding/json"
	"log"
	"my-api/internal/models"
	"my-api/internal/services"
	"net/http"
)

func MakeDecisionHandler(w http.ResponseWriter, r *http.Request) {

	var req models.MakeDecisionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erreur de décodage du JSON", http.StatusBadRequest)
		return
	}

	back, err := services.MakeDecisionService(req.Message)

	if err != nil {
		http.Error(w, "Erreur using chatgpt api", http.StatusBadRequest)
		return
	}

	res := models.MakeDecisionResponse{
		Message: back,
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
