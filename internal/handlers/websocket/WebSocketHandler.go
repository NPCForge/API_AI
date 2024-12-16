package websocketHandlers

import (
	"encoding/json"
	"log"
	"my-api/internal/models/websocket"
	"my-api/internal/services/websocket"

	"my-api/internal/utils"

	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow CORS
	},
}

var actions = []websocketModels.WebSocketDispatcher{
	{
		Name:      "Register",
		Handler:   RegisterHandlerWebsocket,
		Protected: false,
	},
	{
		Name:      "Connection",
		Handler:   ConnectHandlerWebSocket,
		Protected: false,
	},
	{
		Name:      "Disconnect",
		Handler:   DisconnectHandlerWebsocket,
		Protected: true,
	},
	{
		Name:      "TakeDecision",
		Handler:   MakeDecisionHandlerWebSocket,
		Protected: true,
	},
}

func handleWebSocketMessage(conn *websocket.Conn, messageType int, message []byte) {
	var msg websocketModels.WebSocketMessage
	var initialRoute = "root"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		utils.SendError(conn, "Error while decoding JSON message", initialRoute)
		return
	}

	if msg.Action == "" {
		utils.SendError(conn, "Missing required fields in the JSON message", initialRoute)
		return
	}

	// Handle action
	for _, action := range actions {
		if action.Name == msg.Action {
			if action.Protected {
				// Call login middleware
				if !websocketServices.LoginMiddlewareWebSocket(conn, message, utils.SendResponse, utils.SendError) {
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
