package httpHandlers

import (
	"encoding/json"
	"log"
	httpModels "my-api/internal/models/http"
	httpServices "my-api/internal/services/http"
	"net/http"
)

func GetPopulationHandler(w http.ResponseWriter, r *http.Request) {
	store := httpServices.GetPopulation()

	res := httpModels.ResponseGetPopulation{
		Data:   store,
		Status: 200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
