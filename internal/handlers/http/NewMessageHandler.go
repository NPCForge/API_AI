package httpHandlers

import (
	"encoding/json"
	"log"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"net/http"
)

func NewMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unauthorized method", http.StatusMethodNotAllowed)
		return
	}

	var req sharedModel.NewMessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if req.Sender == "" || req.Receivers == nil || req.Message == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = sharedServices.NewMessageService(req.Sender, req.Receivers, req.Message)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := sharedModel.NewMessageResponse{
		Message: "Created",
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error while sending json : %v", err)
	}
}
