package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

// NewMessageHandlerWebSocket handles WebSocket requests to create and send a new message between entities.
func NewMessageHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var req sharedModel.NewMessageRequest
	var initialRoute = "NewMessage"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if req.Sender == "" || req.Receivers == nil || req.Message == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the JSON body",
		})
		return
	}

	err = sharedServices.NewMessageService(req.Sender, req.Receivers, req.Message)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"message": "Success",
	})
}
