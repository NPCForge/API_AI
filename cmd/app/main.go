package main

import (
	"log"
	"my-api/config"
	"my-api/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Crée un nouveau routeur
	r := mux.NewRouter()
	config.InitClient()

	// Associe le handler à la route /internal/handlers/RouteHandler.go
	r.HandleFunc("/Connect", handlers.ConnectHandler).Methods("POST")
	r.HandleFunc("/Register", handlers.RegisterHandler).Methods("POST")

	// Lance le serveur
	log.Println("Serveur démarré sur le port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
	}
}
