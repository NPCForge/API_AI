package main

import (
	"log"
	"my-api/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Crée un nouveau routeur
	r := mux.NewRouter()

	// Associe le handler `ConnectHandler` à la route `/connect`
	r.HandleFunc("/connect", handlers.ConnectHandler).Methods("POST")

	// Lance le serveur
	log.Println("Serveur démarré sur le port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
	}
}
