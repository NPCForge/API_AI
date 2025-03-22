package main

import (
	"log"
	"my-api/config"
	"my-api/pkg"
	"net/http"

	httpHandlers "my-api/internal/handlers/http"
	websocketHandlers "my-api/internal/handlers/websocket"
	httpServices "my-api/internal/services/http"

	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.Lshortfile)
	r := mux.NewRouter()
	config.DrawLogo()
	config.InitDB()
	port := config.ManageArgument()

	// Goroutine pour les commandes
	go pkg.Commande()

	// Websocket handler
	r.HandleFunc("/ws", websocketHandlers.WebsocketHandler).Methods("GET")

	// Http handler
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(httpServices.LoggingMiddleware)

	r.HandleFunc("/Connect", httpHandlers.ConnectHandler).Methods("POST")
	r.HandleFunc("/Register", httpHandlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/Health", httpHandlers.HealthHandler).Methods("GET")

	protected.HandleFunc("/Disconnect", httpHandlers.DisconnectHandler).Methods("POST")
	protected.HandleFunc("/Remove", httpHandlers.RemoveHandler).Methods("POST")
	protected.HandleFunc("/MakeDecision", httpHandlers.MakeDecisionHandler).Methods("POST")
	protected.HandleFunc("/GetPopulation", httpHandlers.GetPopulationHandler).Methods("GET")

	log.Printf("Serveur démarré sur http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
	}
}
