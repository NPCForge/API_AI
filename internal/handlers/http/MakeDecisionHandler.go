package httpHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"net/http"
	"strconv"
)

// MakeDecisionHandler handles POST requests where an entity makes a decision based on a message.
func MakeDecisionHandler(w http.ResponseWriter, r *http.Request) {
	var req sharedModel.MakeDecisionRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if req.Checksum == "" {
		http.Error(w, "Missing required fields in the JSON body", http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")

	data, err := sharedServices.MakeDecisionService(req.Message, req.Checksum, token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data["Status"] = strconv.Itoa(200)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
