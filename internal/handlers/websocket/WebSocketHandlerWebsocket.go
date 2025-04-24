package websocketHandlers

import (
	"encoding/json"
	"log"
	"my-api/internal/handlers"
	websocketModels "my-api/internal/models/websocket"
	websocketServices "my-api/internal/services/websocket"

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
		Name:       "Register",
		Handler:    RegisterHandlerWebsocket,
		Protected:  false,
		IsBlocking: false,
	},
	{
		Name:       "Connect",
		Handler:    ConnectHandlerWebSocket,
		Protected:  false,
		IsBlocking: false,
	},
	{
		Name:       "Disconnect",
		Handler:    DisconnectHandlerWebsocket,
		Protected:  true,
		IsBlocking: false,
	},
	{
		Name:       "TakeDecision",
		Handler:    MakeDecisionHandlerWebSocket,
		Protected:  true,
		IsBlocking: false,
	},
	{
		Name:       "Remove",
		Handler:    RemoveHandlerWebSocket,
		Protected:  true,
		IsBlocking: false,
	},
	{
		Name:       "NewMessage",
		Handler:    NewMessageHandlerWebsocket,
		Protected:  true,
		IsBlocking: false,
	},
	{
		Name:       "ResetGame",
		Handler:    ResetGameWebsocket,
		Protected:  false,
		IsBlocking: true,
	},
	{
		Name:       "CreateEntity",
		Handler:    CreateEntityHandlerWebSocket,
		Protected:  true,
		IsBlocking: false,
	},
	{
		Name:       "RemoveEntity",
		Handler:    RemoveEntityHandlerWebSocket,
		Protected:  true,
		IsBlocking: false,
	},
	{
		Name:       "GetEntities",
		Handler:    GetEntitiesHandlerWebSocket,
		Protected:  true,
		IsBlocking: false,
	},
}

func handleBlockingEvent(triggerConn *websocket.Conn) {
	handlers.WS.BlockMutex.Lock()
	handlers.WS.IsBlocking = true
	handlers.WS.BlockMutex.Unlock()
}

func handleWebSocketMessage(conn *websocket.Conn, messageType int, message []byte) {
	handlers.WS.BlockMutex.RLock()
	if handlers.WS.IsBlocking {
		handlers.WS.BlockMutex.RUnlock()
		utils.SendError(conn, "root", map[string]interface{}{
			"message": "API temporarily unavailable. Please try again later.",
		})
		return
	}
	handlers.WS.BlockMutex.RUnlock()

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
			if action.IsBlocking {
				handleBlockingEvent(conn)
			}
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

	log.Println("New WebSocket connection")

	handlers.WS.Mutex.Lock()
	handlers.WS.Conns[conn] = true
	handlers.WS.Mutex.Unlock()

	defer func() {
		handlers.WS.Mutex.Lock()
		delete(handlers.WS.Conns, conn)
		handlers.WS.Mutex.Unlock()
		err := conn.Close()
		if err != nil {
			return
		}
		log.Println("WebSocket disconnected")
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading :", err)
			break
		}

		handleWebSocketMessage(conn, messageType, message)
	}
}
