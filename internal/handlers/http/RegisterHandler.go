package httpHandlers

import (
	"encoding/json"
	"my-api/config"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"my-api/pkg"
	"net/http"
)

// RegisterHandler handles POST requests to register a new user after verifying an API key.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req sharedModel.RegisterRequest
	var res sharedModel.RegisterResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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

	if req.Password != "" && req.Identifier != "" && req.GamePrompt != "" {
		private, id, err := sharedServices.RegisterService(req.Password, req.Identifier, req.GamePrompt)

		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		res.Private = private
		res.Status = http.StatusCreated
		res.Message = "User created"
		res.Id = id
	} else {
		res.Private = "null"
		res.Status = http.StatusNoContent
		res.Message = "Password or Identifier is missing"
		res.Id = "null"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		pkg.DisplayContext("Error while sending JSON response", pkg.Error, err)
	}
}
