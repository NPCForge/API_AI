package httpHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	service "my-api/internal/services/merged"
	"net/http"
)

func MakeDecisionHandler(w http.ResponseWriter, r *http.Request) {
	var req sharedModel.MakeDecisionRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if req.Message == "" || req.Checksum == "" {
		http.Error(w, "Missing required fields in the JSON message", http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")

	msg, err := service.MakeDecisionService(req.Message, req.Checksum, token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := sharedModel.MakeDecisionResponse{
		Message: msg,
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}
