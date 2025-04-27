package websocketHandlers

import (
	"encoding/json"
	"log"
	websocketModels "my-api/internal/models/websocket"
	websocketServices "my-api/internal/services/websocket"
	"my-api/pkg"

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
		Name:      "Connect",
		Handler:   ConnectHandlerWebSocket,
		Protected: false,
	},
	{
		Name:      "Disconnect",
		Handler:   DisconnectHandlerWebsocket,
		Protected: true,
	},
	{
		Name:      "MakeDecision",
		Handler:   MakeDecisionHandlerWebSocket,
		Protected: true,
	},
	{
		Name:      "RemoveUser",
		Handler:   RemoveUserHandlerWebSocket,
		Protected: true,
	},
	{
		Name:      "NewMessage",
		Handler:   NewMessageHandlerWebSocket,
		Protected: true,
	},
	{
		Name:      "ResetGame",
		Handler:   ResetGameWebsocket,
		Protected: false,
	},
	{
		Name:      "CreateEntity",
		Handler:   CreateEntityHandlerWebSocket,
		Protected: true,
	},
	{
		Name:      "RemoveEntity",
		Handler:   RemoveEntityHandlerWebSocket,
		Protected: true,
	},
	{
		Name:      "GetEntities",
		Handler:   GetEntitiesHandlerWebSocket,
		Protected: true,
	},
}

func handleWebSocketMessage(conn *websocket.Conn, messageType int, message []byte) {
	var msg websocketModels.WebSocketMessage
	var initialRoute = "root"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		utils.SendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if msg.Action == "" {
		utils.SendError(conn, initialRoute, map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	println("Message received: " + msg.Action)

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
	pkg.DisplayContext("Cannot find matching route for: "+msg.Action, pkg.Error)
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading :", err)
		return
	}

	log.Println("New WebSocket connection")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading :", err)
			break
		}

		handleWebSocketMessage(conn, messageType, message)
	}
}
