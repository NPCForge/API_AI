package httpHandlers

import (
	"encoding/json"
	"log"
	httpModels "my-api/internal/models/http"
	service "my-api/internal/services/merged"
	"my-api/pkg"
	"net/http"
)

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	var req httpModels.RemoveRequestRefacto
	var res httpModels.RemoveResponseRefacto

	// delete

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "JSON decoding error", http.StatusBadRequest)
		return
	}

	if req.DeleteUserIdentifier == "" || req.Token == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	disconnect, err := service.RemoveService(req.Token, req.DeleteUserIdentifier)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if disconnect {
		pkg.DisplayContext("Remove Handler need to Disconnect HTTP", pkg.Debug)
	}

	// disconnect

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
