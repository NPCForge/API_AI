package websocketHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	websocketModels "my-api/internal/models/websocket"
	service "my-api/internal/services/merged"
)

func ConnectHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, map[string]interface{}),
) {
	var msg websocketModels.ConnectRequestRefacto
	var initialRoute = "Connect"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if msg.Action == "" || msg.Identifier == "" || msg.Password == "" {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	private, err := service.ConnectService(msg.Password, msg.Identifier)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	sendResponse(conn, initialRoute, map[string]interface{}{
		"token": private,
	})
}
