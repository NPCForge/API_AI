package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	service "my-api/internal/services/merged"

	"github.com/gorilla/websocket"
)

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

	if err := service.RemoveEntityService(req.Checksum, req.Token); err != nil {
		sendError(conn, route, "", map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	sendResponse(conn, route, "", map[string]interface{}{
		"message": "Entity successfully deleted",
	})
}
