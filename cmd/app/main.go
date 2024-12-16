package main

import (
	"fmt"
	"log"

	"my-api/config"

	"my-api/internal/handlers/http"
	"my-api/internal/handlers/websocket"

	"my-api/internal/services/http"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	config.InitDB()

	r.HandleFunc("/ws", websocketHandlers.WebsocketHandler).Methods("GET")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(httpServices.LoggingMiddleware)

	r.HandleFunc("/Connect", httpHandlers.ConnectHandler).Methods("POST")
	r.HandleFunc("/Register", httpHandlers.RegisterHandler).Methods("POST")

	protected.HandleFunc("/Disconnect", httpHandlers.DisconnectHandler).Methods("POST")
	protected.HandleFunc("/Remove", httpHandlers.RemoveHandler).Methods("POST")
	protected.HandleFunc("/MakeDecision", httpHandlers.MakeDecisionHandler).Methods("POST")

	port := ":3000"
	fmt.Printf("Serveur démarré sur http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
	}
}
