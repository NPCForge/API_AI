package httpHandlers

import (
	"encoding/json"
	"my-api/config"
	sharedModel "my-api/internal/models/shared"
	service "my-api/internal/services/merged"
	"my-api/pkg"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req sharedModel.RegisterRequest
	var res sharedModel.RegisterResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	apiKey := config.GetEnvVariable("API_KEY_REGISTER")

	if req.Token != apiKey {
		http.Error(w, "Unauthorized method", http.StatusUnauthorized)
		return
	}

	if req.Password != "" && req.Identifier != "" {
		private, id, err := service.RegisterService(req.Password, req.Identifier)

		if err != nil {
			http.Error(w, "Server Error", 500)
			return
		}

		res.Private = private
		res.Status = 201
		res.Message = "User created"
		res.Id = id
	} else {
		res.Private = "null"
		res.Status = 204
		res.Message = "Password or Identifier is not given"
		res.Id = "null"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		pkg.DisplayContext("Error while sending JSON response", pkg.Error, err)
	}
}
