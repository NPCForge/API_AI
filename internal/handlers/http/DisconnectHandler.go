package httpHandlers

import (
	"encoding/json"
	"log"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"net/http"
)

// DisconnectHandler handles user disconnection requests.
func DisconnectHandler(w http.ResponseWriter, r *http.Request) {
	res := sharedModel.DisconnectResponse{
		Message: "Successfully disconnected",
		Status:  200,
	}

	token := r.Header.Get("Authorization")
	err := sharedServices.DisconnectService(token)

	if err != nil {
		res = sharedModel.DisconnectResponse{
			Message: "Error while disconnecting",
			Status:  401,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error while sending JSON: %v", err)
	}
}
