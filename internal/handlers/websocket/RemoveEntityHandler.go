package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

// RemoveEntityHandlerWebSocket handles WebSocket requests to delete an entity by its checksum.
func RemoveEntityHandlerWebSocket(
	conn *websocket.Conn,
	message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	const route = "RemoveEntity"

	var req sharedModel.RemoveEntityRequest
	if err := json.Unmarshal(message, &req); err != nil {
		sendError(conn, route, "", map[string]interface{}{
			"message": "Error decoding JSON payload",
		})
		return
	}

	if err := sharedServices.RemoveEntityService(req.Checksum, req.Token); err != nil {
		sendError(conn, route, "", map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	sendResponse(conn, route, "", map[string]interface{}{
		"message": "Entity successfully deleted",
	})
}
