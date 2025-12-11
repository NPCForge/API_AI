package websocketHandlers

import (
	"encoding/json"
	"log"
	websocketModels "my-api/internal/models/websocket"
	websocketServices "my-api/internal/services/websocket"
	"my-api/internal/utils"
	"my-api/pkg"

	"net/http"

	"my-api/internal/services"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	ActiveClients      = make(map[*websocket.Conn]int)
	ActiveClientsMutex sync.Mutex
)

func RegisterClient(conn *websocket.Conn, userID int) {
	ActiveClientsMutex.Lock()
	defer ActiveClientsMutex.Unlock()
	ActiveClients[conn] = userID
}

func UnregisterClient(conn *websocket.Conn) int {
	ActiveClientsMutex.Lock()
	defer ActiveClientsMutex.Unlock()
	userID, exists := ActiveClients[conn]
	if exists {
		delete(ActiveClients, conn)
		return userID
	}
	return 0
}

// WebSocket upgrader configuration allowing CORS.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

// actions defines all available WebSocket routes and their handlers.
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
		Name:      "Restart",
		Handler:   RestartHandlerWebSocket,
		Protected: true,
	},
	{
		Name:      "GetEntities",
		Handler:   GetEntitiesHandlerWebSocket,
		Protected: true,
	},
	{
		Name:      "Status",
		Handler:   StatusHandlerWebSocket,
		Protected: true,
	},
}

// handleWebSocketMessage dispatches an incoming WebSocket message to the appropriate handler based on the action field.
func handleWebSocketMessage(conn *websocket.Conn, messageType int, message []byte) {
	var msg websocketModels.WebSocketMessage
	var initialRoute = "root"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		utils.SendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if msg.Action == "" {
		utils.SendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the JSON body",
		})
		return
	}

	println("Message received: " + msg.Action)

	// Handle the action
	for _, action := range actions {
		if action.Name == msg.Action {
			if action.Protected {
				// Check authentication with middleware
				if !websocketServices.LoginMiddlewareWebSocket(conn, message, utils.SendResponse, utils.SendError) {
					return
				}
			}
			// Call the corresponding action handler
			action.Handler(conn, message, utils.SendResponse, utils.SendError)
			return
		}
	}
	pkg.DisplayContext("Cannot find matching route for: "+msg.Action, pkg.Error)
}

// WebsocketHandler upgrades HTTP requests to WebSocket connections and listens for incoming messages.
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading:", err)
		return
	}

	log.Println("New WebSocket connection established")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading:", err)

			// Cleanup user on disconnect
			userID := UnregisterClient(conn)
			if userID != 0 {
				log.Printf("Cleaning up user ID: %d", userID)
				if err := services.DropUser(userID); err != nil {
					log.Printf("Error dropping user %d: %v", userID, err)
				}
			}

			break
		}

		handleWebSocketMessage(conn, messageType, message)
	}
}
