package httpHandlers

import (
	"encoding/json"
	"fmt"
	sharedModel "my-api/internal/models/shared"
	"net/http"

	sharedServices "my-api/internal/services/shared"
)

// CreateEntityHandler handles POST requests to create a new entity.
func CreateEntityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req sharedModel.RequestCreateEntity
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.Checksum == "" || req.Name == "" || req.Prompt == "" || req.Role == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(sharedModel.ResponseCreateEntity{
			Id:      "",
			Message: "Invalid request body",
			Status:  "error",
		})
		return
	}

	token := r.Header.Get("Authorization")

	id, err := sharedServices.CreateEntityService(req.Name, req.Prompt, req.Checksum, token, req.Role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(sharedModel.ResponseCreateEntity{
			Id:      "",
			Message: "Failed to create entity",
			Status:  "error",
		})
		return
	}

	resp := sharedModel.ResponseCreateEntity{
		Id:       fmt.Sprintf("%d", id),
		Message:  "Entity created successfully",
		Checksum: req.Checksum,
		Status:   "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
