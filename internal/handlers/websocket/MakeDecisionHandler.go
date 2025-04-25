package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	service "my-api/internal/services/merged"

	"github.com/gorilla/websocket"
)

func MakeDecisionHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, map[string]interface{}),
) {
	var req sharedModel.MakeDecisionRequest
	var initialRoute = "MakeDecision"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if req.Message == "" || req.Checksum == "" {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	msg, err := service.MakeDecisionService(req.Message, req.Checksum, req.Token)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": err,
		})
		return
	}

	sendResponse(conn, initialRoute, map[string]interface{}{
		"message": msg,
	})
}
