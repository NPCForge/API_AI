package httpHandlers

import (
	"encoding/json"
	"log"
	sharedModel "my-api/internal/models/shared"
	service "my-api/internal/services/merged"
	"net/http"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unauthorized method", http.StatusMethodNotAllowed)
		return
	}

	var req sharedModel.ConnectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	pass, id, err := service.ConnectService(req.Password, req.Identifier)
	var res sharedModel.ConnectResponse

	if err != nil {
		res = sharedModel.ConnectResponse{
			Message:  "Unauthorized",
			Status:   401,
			TmpToken: "",
		}
	} else {
		res = sharedModel.ConnectResponse{
			Message:  "Successfully connected",
			Status:   200,
			Id:       id,
			TmpToken: pass,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error while sending json : %v", err)
	}
}
