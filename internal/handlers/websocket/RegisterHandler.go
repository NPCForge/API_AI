package websocketHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"my-api/config"
	"my-api/internal/models/websocket"
	"my-api/internal/services/websocket"
)

func RegisterHandlerWebsocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg websocketModels.RegisterRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

	if msg.Action == "" || msg.Token == "" || msg.Checksum == "" || msg.Name == "" || msg.Prompt == "" {
		sendError(conn, "Missing required fields in the JSON message")
		return
	}

	apiKey := config.GetEnvVariable("API_KEY_REGISTER")

	// Check if the token is valid
	if msg.Token != apiKey {
		sendError(conn, "Invalid API Key")
		return
	}

	websocketServices.RegisterServiceWebSocket(conn, msg, sendResponse, sendError)
}
