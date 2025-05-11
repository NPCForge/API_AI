package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"my-api/pkg"

	"github.com/gorilla/websocket"
)

// MakeDecisionHandlerWebSocket handles WebSocket requests where an entity makes a decision based on a message.
func MakeDecisionHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var req sharedModel.MakeDecisionRequest
	var initialRoute = "MakeDecision"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, req.Checksum, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if req.Checksum == "" {
		sendError(conn, initialRoute, req.Checksum, map[string]interface{}{
			"message": "Missing required fields in the JSON body",
		})
		return
	}

	msg, err := sharedServices.MakeDecisionService(req.Message, req.Checksum, req.Token)
	if err != nil {
		pkg.DisplayContext("Internal server error", pkg.Error, err)
		sendError(conn, initialRoute, req.Checksum, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	sendResponse(conn, initialRoute, req.Checksum, map[string]interface{}{
		"message": msg,
	})
}
