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
)

func main() {
	log.SetFlags(log.Lshortfile)
	r := mux.NewRouter()
	config.DrawLogo()
	config.InitDB()

	// Goroutine pour les commandes
	if !utils.IsRunningInDocker() {
		go utils.Commande()
	}

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
	pkg.DisplayContext(fmt.Sprintf("Serveur démarré sur http://localhost%s", port), pkg.Update)
	if err := http.ListenAndServe(port, r); err != nil {
		pkg.DisplayContext("Erreur lors du lancement du serveur", pkg.Error, err, true)
	}
}
