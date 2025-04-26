package httpHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	service "my-api/internal/services/merged"
	"my-api/pkg"
	"net/http"
)

func RemoveEntityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unauthorized method", http.StatusMethodNotAllowed)
		return
	}

	var req sharedModel.RemoveEntityRequest
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

	err = service.RemoveEntityService(req.Checksum, token)

	var res sharedModel.RemoveEntityResponse

	if err != nil {
		pkg.DisplayContext("Error during RemoveEntityService: ", pkg.Error, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		res = sharedModel.RemoveEntityResponse{
			Message: "Successfully deleted",
			Status:  200,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
