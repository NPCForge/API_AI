package websocketHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
)

// DisconnectHandlerWebsocket handles WebSocket requests to disconnect a user.
func DisconnectHandlerWebsocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var msg sharedModel.DisconnectRequest
	var initialRoute = "Disconnect"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if msg.Action == "" || msg.Token == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the JSON body",
		})
		return
	}

	err = sharedServices.DisconnectService(msg.Token)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while disconnecting",
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"message": "Successfully disconnected",
	})
}
