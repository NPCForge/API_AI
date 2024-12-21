package httpHandlers

import (
	"encoding/json"
	"log"
	http3 "my-api/internal/models/http"
	http2 "my-api/internal/services/http"
	"net/http"
)

func MakeDecisionHandler(w http.ResponseWriter, r *http.Request) {

	var req http3.MakeDecisionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erreur de décodage du JSON", http.StatusBadRequest)
		return
	}

	back, err := http2.MakeDecisionService(req.Message)

	if err != nil {
		http.Error(w, "Erreur using chatgpt api", http.StatusBadRequest)
		return
	}

	res := http3.MakeDecisionResponse{
		Message: back,
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erreur lors de l'envoi de la réponse JSON : %v", err)
	}
}
