package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

func ConnectHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var msg sharedModel.ConnectRequest
	var initialRoute = "Connect"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if msg.Action == "" || msg.Identifier == "" || msg.Password == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	private, id, err := sharedServices.ConnectService(msg.Password, msg.Identifier)

	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"token": private,
		"id":    id,
	})
}
