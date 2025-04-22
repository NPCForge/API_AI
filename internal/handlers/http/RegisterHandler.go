package httpHandlers

import (
	"encoding/json"
	"my-api/config"
	httpModels "my-api/internal/models/http"
	service "my-api/internal/services/merged"
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

	if req.Password != "" && req.Identifiant != "" {
		private, err := service.RegisterService(req.Identifiant, req.Password)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		res.Private = private
		res.Status = 201
	}

	// private, err := httpServices.RegisterService(req)

	// Check if the token was successfully created
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", 500)
	// 	return
	// }

	// // Create the response
	// res := httpModels.RegisterResponse{
	// 	Message: "Connection successful",
	// 	Status:  200,
	// 	Private: private,
	// }

	// // Set the Content-Type header to application/json and send the JSON response
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(res); err != nil {
	// 	log.Printf("Error while sending JSON response: %v", err)
	// }
}
