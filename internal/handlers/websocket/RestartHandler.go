package websocketHandlers

import (
	"encoding/json"

	websocketModels "my-api/internal/models/websocket"
	websocketServices "my-api/internal/services/websocket"

	"github.com/gorilla/websocket"
)

// RestartHandlerWebSocket processes a restart request over WebSocket.
func RestartHandlerWebSocket(
	conn *websocket.Conn,
	message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var msg websocketModels.RestartRequest
	initialRoute := "Restart"

	if err := json.Unmarshal(message, &msg); err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{"message": "Error while decoding JSON message"})
		return
	}

	if msg.Token == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{"message": "Missing token"})
		return
	}

	websocketServices.RestartServiceWebSocket(conn, msg.Token, sendResponse, sendError)
}