package websocketHandlers

import (
	"encoding/json"
	"my-api/internal/models/websocket"
	"my-api/internal/services/websocket"

	"github.com/gorilla/websocket"
)

func MakeDecisionHandlerWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, string, string), sendError func(*websocket.Conn, string, string)) {
	var msg websocketModels.MakeDecisionRequest
	var initialRoute = "MakeDecision"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message", initialRoute)
		return
	}

	if msg.Message == "" {
		sendError(conn, "Missing required fields in the JSON message", initialRoute)
		return
	}

	websocketServices.MakeDecisionWebSocket(conn, msg, sendResponse, sendError)
}
