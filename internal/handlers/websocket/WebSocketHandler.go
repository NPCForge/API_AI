package websocket

import (
	"encoding/json"
	"log"

	"my-api/internal/services"
	"my-api/internal/utils"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow CORS
	},
}

func handleWebSocketMessage(conn *websocket.Conn, messageType int, message []byte) {
	// Log hard message
	log.Printf("Message : %s", message)

	// Decode Json message
	var msg struct {
		Action  string `json:"action"`
	}

	err := json.Unmarshal(message, &msg)
	if err != nil {
		utils.SendError(conn, "Error while decoding JSON message")
		return
	}

	// Log decoded messages
	log.Printf("Action : %s", msg.Action)

	// Traite l'action
	switch msg.Action {
        case "TakeDecision":
            services.MakeDecisionWebSocket(conn, message, utils.SendResponse, utils.SendError)
        default:
            utils.SendError(conn, "Unknown Action")
	}
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading :", err)
		return
	}
	defer conn.Close()

	log.Println("New WebSocket connection")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading :", err)
			break
		}

		// Handle websockets messages
		handleWebSocketMessage(conn, messageType, message)
	}
}