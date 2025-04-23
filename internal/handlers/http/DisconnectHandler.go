package httpHandlers

import (
	"encoding/json"
	"log"
	httpModels "my-api/internal/models/http"
	httpServices "my-api/internal/services/http"
	"net/http"
)

func DisconnectHandler(w http.ResponseWriter, r *http.Request) {
	res := httpModels.DisconnectResponse{
		Message: "Deconnexion réussie",
		Status:  200,
	}

	token := r.Header.Get("Authorization")
	_, err := httpServices.Disconnect(token)

	if err != nil {
		res = httpModels.DisconnectResponse{
			Message: "Erreur lors de la deconnexion",
			Status:  401,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
