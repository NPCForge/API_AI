package main

import (
	"fmt"
	"log"
	"my-api/config"
	"my-api/pkg"
	"net/http"

	httpHandlers "my-api/internal/handlers/http"
	websocketHandlers "my-api/internal/handlers/websocket"
	httpServices "my-api/internal/services/http"
	"my-api/internal/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors" // Importer le package CORS
)

// Health responds to a health check request with "OK".
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

// main initializes the server, database, and routes, then starts the HTTP server.
func main() {
	log.SetFlags(log.Lshortfile)

	r := mux.NewRouter()

	// Display the server logo
	config.DrawLogo()

	// Initialize the database connection
	config.InitDB()

	// Launches a goroutine for CLI commands if not running in Docker
	if !utils.IsRunningInDocker() {
		go utils.Commande()
	}

	// WebSocket handler
	r.HandleFunc("/ws", websocketHandlers.WebsocketHandler).Methods("GET")

	// HTTP handlers
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(httpServices.LoggingMiddleware)

	// Public routes
	r.HandleFunc("/Connect", httpHandlers.ConnectHandler).Methods("POST")
	r.HandleFunc("/Register", httpHandlers.RegisterHandler).Methods("POST")

	// Protected routes
	protected.HandleFunc("/Disconnect", httpHandlers.DisconnectHandler).Methods("POST")
	protected.HandleFunc("/RemoveUser", httpHandlers.RemoveUserHandler).Methods("POST")
	protected.HandleFunc("/MakeDecision", httpHandlers.MakeDecisionHandler).Methods("POST")
	protected.HandleFunc("/CreateEntity", httpHandlers.CreateEntityHandler).Methods("POST")
	protected.HandleFunc("/RemoveEntity", httpHandlers.RemoveEntityHandler).Methods("POST")
	protected.HandleFunc("/NewMessage", httpHandlers.NewMessageHandler).Methods("POST")
	protected.HandleFunc("/GetEntities", httpHandlers.GetEntitiesHandler).Methods("GET")

	// Health check route
	r.HandleFunc("/health", Health).Methods("GET")

	// CORS handler - Permet de gérer les requêtes depuis localhost:8001
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8001"}, // Ajouter ton origine front-end
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Appliquer CORS aux routes
	handler := c.Handler(r)

	// Start the server
	port := "0.0.0.0:3000"
	pkg.DisplayContext(fmt.Sprintf("Server started on http://%s", port), pkg.Update)
	if err := http.ListenAndServe(port, handler); err != nil {
		pkg.DisplayContext("Error while starting the server", pkg.Error, err, true)
	}
}
