package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

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
			"message": "Error while getting entities",
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"entities": entities,
	})
}
