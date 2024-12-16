package websocketServices

import (
	"my-api/internal/models/websocket"
	"my-api/internal/services"

	"github.com/gorilla/websocket"
)

func MakeDecisionWebSocket(conn *websocket.Conn, msg websocketModels.MakeDecisionRequest, sendResponse func(*websocket.Conn, string, string), sendError func(*websocket.Conn, string, string)) {
	var initialRoute = "MakeDecision"

	back, err := services.GptSimpleRequest(msg.Message)
	if err != nil {
		sendError(conn, "Error while calling MakeDecision service", initialRoute)
		return
	}

	sendResponse(conn, back, initialRoute)
}
