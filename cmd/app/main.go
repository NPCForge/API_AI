package main

import (
	"log"
	"my-api/config"
	"net/http"

	httpHandlers "my-api/internal/handlers/http"
	websocketHandlers "my-api/internal/handlers/websocket"
	httpServices "my-api/internal/services/http"
	"my-api/internal/utils"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	config.DrawLogo()
	config.InitDB()

	go utils.Commande()

	// Websocket handler
	r.HandleFunc("/ws", websocketHandlers.WebsocketHandler).Methods("GET")

	// Http handler
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(httpServices.LoggingMiddleware)

	r.HandleFunc("/Connect", httpHandlers.ConnectHandler).Methods("POST")
	r.HandleFunc("/Register", httpHandlers.RegisterHandler).Methods("POST")

	protected.HandleFunc("/Disconnect", httpHandlers.DisconnectHandler).Methods("POST")
	protected.HandleFunc("/Remove", httpHandlers.RemoveHandler).Methods("POST")
	protected.HandleFunc("/MakeDecision", httpHandlers.MakeDecisionHandler).Methods("POST")

	port := ":3000"
	log.Printf("Serveur démarré sur http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
	}
}
