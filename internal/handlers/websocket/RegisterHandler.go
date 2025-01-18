package websocketHandlers

import (
	"encoding/json"
	"log"
	"my-api/config"
	websocketModels "my-api/internal/models/websocket"
	websocketServices "my-api/internal/services/websocket"

	"github.com/gorilla/websocket"
)

func RegisterHandlerWebsocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	log.Println("RegisterHandlerWebsocket")
	var msg websocketModels.RegisterRequest
	var initialRoute = "Register"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if msg.Action == "" || msg.Token == "" || msg.Checksum == "" || msg.Name == "" || msg.Prompt == "" {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	apiKey := config.GetEnvVariable("API_KEY_REGISTER")

	// Check if the token is valid
	if msg.Token != apiKey {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Invalid API Key",
		})
		return
	}

	websocketServices.RegisterServiceWebSocket(conn, msg, sendResponse, sendError)
}
