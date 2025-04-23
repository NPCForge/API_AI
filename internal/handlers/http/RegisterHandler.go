package httpHandlers

import (
	"encoding/json"
	"my-api/config"
	httpModels "my-api/internal/models/http"
	service "my-api/internal/services/merged"
	"my-api/pkg"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req httpModels.RegisterRequestRefacto
	var res httpModels.RegisterResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "JSON decoding error", http.StatusBadRequest)
		return
	}

	apiKey := config.GetEnvVariable("API_KEY_REGISTER")

	// Check if the token is valid
	if req.Token != apiKey {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if req.Password != "" && req.Identifier != "" {
		private, err := service.RegisterService(req.Password, req.Identifier)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		res.Private = private
		res.Status = 201
		res.Message = "User created"
	} else {
		res.Private = "null"
		res.Status = 204
		res.Message = "Password or Identifier is not given"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		pkg.DisplayContext("Error while sending JSON response", pkg.Error, err)
	}
}
