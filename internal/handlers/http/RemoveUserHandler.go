package httpHandlers

import (
	"encoding/json"
	"log"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"net/http"
)

// RemoveUserHandler handles POST requests to remove a user by username.
func RemoveUserHandler(w http.ResponseWriter, r *http.Request) {
	var req sharedModel.RemoveUserRequest
	var res sharedModel.RemoveUserResponse

	res = sharedModel.RemoveUserResponse{
		Message: "Successfully deleted",
		Status:  200,
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON Decoding Error", http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")

	err = sharedServices.RemoveUserService(token, req.UserName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error while sending JSON response: %v", err)
	}
}
