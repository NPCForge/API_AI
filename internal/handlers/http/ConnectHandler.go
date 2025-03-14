package httpHandlers

import (
	"encoding/json"
	"log"
	httpModels "my-api/internal/models/http"
	httpServices "my-api/internal/services/http"
	"net/http"
)

// ConnectHandler gère la requête pour la route Connect
func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	// Vérifie que la méthode est POST
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Décode le corps de la requête
	var req httpModels.ConnectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erreur de décodage du JSON", http.StatusBadRequest)
		return
	}

	pass, err := httpServices.UserConnect(req)
	var res httpModels.ConnectResponse

	if err != nil {
		res = httpModels.ConnectResponse{
			Message:  "Unauthorized",
			Status:   401,
			TmpToken: "",
		}
	} else {
		res = httpModels.ConnectResponse{
			Message:  "Connexion réussie",
			Status:   200,
			TmpToken: pass,
		}
	}

	// Définit l'en-tête Content-Type à application/json et envoie la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
