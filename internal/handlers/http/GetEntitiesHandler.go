package httpHandlers

import (
	"encoding/json"
	"fmt"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"net/http"
)

func GetEntitiesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetEntitiesHandler called")

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authorization")

	entities, err := sharedServices.GetEntitiesService(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(sharedModel.ResponseGetEntities{
			Entity: nil,
			Status: "error",
		})
		return
	}

	resp := sharedModel.ResponseGetEntities{
		Entity: entities,
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}
