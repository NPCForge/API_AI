package websocketServices

import (
	"my-api/internal/models/websocket"
	"my-api/internal/services"

	"github.com/gorilla/websocket"
)

func MakeDecisionWebSocket(conn *websocket.Conn, msg websocketModels.MakeDecisionRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "MakeDecision"

	back, err := services.GptSimpleRequest(msg.Message)
	if err != nil {
		println("Error in MakeDecisionWebSocket: " + err.Error())
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while calling MakeDecision service",
		})
		return
	}

	sendResponse(conn, initialRoute, map[string]interface{}{
		"message": back,
	})
}
