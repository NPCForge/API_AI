package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

// EntityDiedHandlerWebSocket handles WebSocket requests to notify about an entity death.
func EntityDiedHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var req sharedModel.EntityDiedRequest
	var initialRoute = "EntityDied"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if req.Checksum == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the payload",
		})
		return
	}

	// Helper or Service call
	err = sharedServices.EntityDiedService(req.Checksum, req.Token)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while processing entity death: " + err.Error(),
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"message": "Entity death processed successfully",
		"status":  200,
	})
}
