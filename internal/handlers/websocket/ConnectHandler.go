package websocketHandlers

import (
	"encoding/json"
	"strconv"

	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	websocketServices "my-api/internal/services/websocket"

	"github.com/gorilla/websocket"
)

// ConnectHandlerWebSocket handles WebSocket connection requests for user authentication.
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
			"message": "Missing required fields in the JSON body",
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

	if intID, convErr := strconv.Atoi(id); convErr == nil {
		websocketServices.UnmarkUserResetting(intID)
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"token": private,
		"id":    id,
	})
}
