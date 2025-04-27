package exemples

import (
	"encoding/json"
	"log"
	"net/http"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unauthorized method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// pass, err := NameService()
	var res Response

	if err != nil {
		res = Response{
			Message: "Unauthorized",
			Status:  401,
		}
	} else {
		res = Response{
			Message: "Successfully connected",
			Status:  200,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error while sending json : %v", err)
	}
}
