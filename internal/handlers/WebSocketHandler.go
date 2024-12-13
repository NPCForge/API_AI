package handlers

import (
	"encoding/json"
	"log"

	"my-api/internal/models"
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

var actions = []models.WebSocketDispatcher{
	{
		Name:      "TakeDecision",
		Handler:   services.MakeDecisionWebSocket,
		Protected: true,
	},
	{
		Name:      "Register",
		Handler:   services.RegisterServiceWebSocket,
		Protected: false,
	},
    {
		Name:      "Connection",
		Handler:   services.UserConnectWebSocket,
		Protected: false,
	},
}

func handleWebSocketMessage(conn *websocket.Conn, messageType int, message []byte) {
	var msg struct {
		Action string `json:"action"`
	}

	err := json.Unmarshal(message, &msg)
	if err != nil {
		utils.SendError(conn, "Error while decoding JSON message")
		return
	}

	// Handle action
	for _, action := range actions {
		if action.Name == msg.Action {
			if action.Protected {
				// Call login middleware
				if !services.LoginMiddlewareWebSocket(conn, message, utils.SendResponse, utils.SendError) {
					return
				}
			}
			// Call action handler
			action.Handler(conn, message, utils.SendResponse, utils.SendError)
			return
		}
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
