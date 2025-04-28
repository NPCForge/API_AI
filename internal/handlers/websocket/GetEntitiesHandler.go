package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

// GetEntitiesHandlerWebSocket handles WebSocket requests to retrieve all entities associated with a user token.
func GetEntitiesHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var req sharedModel.RequestGetEntities
	var initialRoute = "GetEntities"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	entities, err := sharedServices.GetEntitiesService(req.Token)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while retrieving entities",
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"entities": entities,
	})
}
