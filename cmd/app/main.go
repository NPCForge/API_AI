package main

import (
	"fmt"
	"log"
	"my-api/config"
	"my-api/internal/handlers"
	"my-api/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Register api key: ", config.GetEnvVariable("API_KEY_REGISTER"))

	r := mux.NewRouter()
	config.InitDB()

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(services.LoggingMiddleware)

	// Associe le handler à la route /internal/handlers/RouteHandler.go
	r.HandleFunc("/Connect", handlers.ConnectHandler).Methods("POST")
	r.HandleFunc("/Register", handlers.RegisterHandler).Methods("POST")

	// Route Protégé par le middleware
	protected.HandleFunc("/Disconnect", handlers.DisconnectHandler).Methods("POST")
	protected.HandleFunc("/Remove", handlers.RemoveHandler).Methods("POST")
	protected.HandleFunc("/MakeDecision", handlers.MakeDecisionHandler).Methods("POST")

	// Lance le serveur
	log.Println("Serveur démarré sur le port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
	}
}
