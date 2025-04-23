package httpHandlers

import (
	"encoding/json"
	"log"
	httpModels "my-api/internal/models/http"
	service "my-api/internal/services/merged"
	"net/http"
)

func DisconnectHandler(w http.ResponseWriter, r *http.Request) {
	res := httpModels.DisconnectResponse{
		Message: "Successfully disconnected",
		Status:  200,
	}

	token := r.Header.Get("Authorization")
	err := service.DisconnectService(token)

	if err != nil {
		res = httpModels.DisconnectResponse{
			Message: "Error while disconnecting",
			Status:  401,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error while sending json : %v", err)
	}
}
