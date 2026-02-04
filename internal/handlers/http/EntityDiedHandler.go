package httpHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"my-api/pkg"
	"net/http"
)

// EntityDiedHandler handles POST requests to notify about an entity death and delete it.
func EntityDiedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unauthorized method", http.StatusMethodNotAllowed)
		return
	}

	var req sharedModel.EntityDiedRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if req.Checksum == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")

	err = sharedServices.EntityDiedService(req.Checksum, token)

	var res sharedModel.EntityDiedResponse

	if err != nil {
		pkg.DisplayContext("Error during EntityDiedService:", pkg.Error, err)
		// We might want to return 500 or 400 depending on error, sticking to 500 for generic service error often used in this codebase
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		res = sharedModel.EntityDiedResponse{
			Message: "Entity death processed successfully",
			Status:  200,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
