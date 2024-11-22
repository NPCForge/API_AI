package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my-api/internal/models"
	"my-api/internal/services"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func MakeDecisionWebSocketHandler(conn *websocket.Conn, message string) {
	// Créer une requête à partir du message
	req := models.MakeDecisionRequest{
		Message: message,
	}

	// Appeler le service MakeDecision
	back, err := services.MakeDecisionService(req.Message)
	if err != nil {
		sendError(conn, "Erreur lors de l'appel au service MakeDecision")
		return
	}

	// Créer la réponse
	res := models.MakeDecisionResponse{
		Message: back,
		Status:  200,
	}

	log.Printf("back = %s", back)

	// Envoyer la réponse via WebSocket
	sendResponse(conn, res)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permet toutes les origines (CORS)
	},
}

func sendResponse(conn *websocket.Conn, response interface{}) {
	resp := map[string]interface{}{
		"status": "success",
		"data":   response,
	}
	conn.WriteJSON(resp)
}

func sendError(conn *websocket.Conn, errorMessage string) {
	resp := map[string]interface{}{
		"status":  "error",
		"message": errorMessage,
	}
	conn.WriteJSON(resp)
}

func handleWebSocketMessage(conn *websocket.Conn, messageType int, message []byte) {
	// Log du message brut reçu
	log.Printf("Message brut reçu : %s", message)

	// Décode le message JSON avec la structure correcte
	var msg struct {
		Action  string `json:"action"`
		Message string `json:"message"`
	}

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Erreur de décodage du message JSON")
		return
	}

	// Log les champs décodés
	log.Printf("Action : %s, Message : %s", msg.Action, msg.Message)

	// Traite l'action
	switch msg.Action {
	case "TakeDecision":
		MakeDecisionWebSocketHandler(conn, msg.Message)
	default:
		sendError(conn, "Action non reconnue")
	}
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur lors de l'upgrade :", err)
		return
	}
	defer conn.Close()

	log.Println("Nouvelle connexion WebSocket établie")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Erreur lors de la lecture :", err)
			break
		}

		// Gestion des messages WebSocket
		handleWebSocketMessage(conn, messageType, message)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ws", websocketHandler).Methods("GET")

	port := ":3000"
	fmt.Printf("Serveur démarré sur http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Erreur du serveur :", err)
	}
}

// func main() {

// 	fmt.Println("Register api key: ", config.GetEnvVariable("API_KEY_REGISTER"))

// 	r := mux.NewRouter()
// 	config.InitDB()

// 	protected := r.PathPrefix("/").Subrouter()
// 	protected.Use(services.LoggingMiddleware)

// 	// Associe le handler à la route /internal/handlers/RouteHandler.go
// 	r.HandleFunc("/Connect", handlers.ConnectHandler).Methods("POST")
// 	r.HandleFunc("/Register", handlers.RegisterHandler).Methods("POST")

// 	// Route Protégé par le middleware
// 	protected.HandleFunc("/Disconnect", handlers.DisconnectHandler).Methods("POST")
// 	protected.HandleFunc("/Remove", handlers.RemoveHandler).Methods("POST")
// 	protected.HandleFunc("/MakeDecision", handlers.MakeDecisionHandler).Methods("POST")

// 	// Lance le serveur
// 	log.Println("Serveur démarré sur le port 8080")
// 	if err := http.ListenAndServe(":8080", r); err != nil {
// 		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
// 	}
// }
