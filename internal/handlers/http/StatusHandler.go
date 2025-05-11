package httpHandlers

import (
	"encoding/json"
	"log"
	sharedModel "my-api/internal/models/shared"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	res := sharedModel.StatusResponse{
		Message: "Authentified",
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error while sending json : %v", err)
	}
}
