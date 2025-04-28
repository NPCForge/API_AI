package websocketHandlers

import (
	"encoding/json"
	"fmt"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

// CreateEntityHandlerWebSocket handles WebSocket requests to create a new entity.
func CreateEntityHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var req sharedModel.RequestCreateEntity
	var initialRoute = "CreateEntity"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if req.Checksum == "" || req.Name == "" || req.Prompt == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the payload",
		})
		return
	}

	id, err := sharedServices.CreateEntityService(req.Name, req.Prompt, req.Checksum, req.Token)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while creating entity",
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"message":  "Success",
		"id":       fmt.Sprintf("%d", id),
		"checksum": req.Checksum,
	})
}
