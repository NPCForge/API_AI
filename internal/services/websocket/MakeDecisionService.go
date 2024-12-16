package websocketServices

import (
	"my-api/internal/models/websocket"
	"my-api/internal/services"

	"github.com/gorilla/websocket"
)

func MakeDecisionWebSocket(conn *websocket.Conn, msg websocketModels.MakeDecisionRequest, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	back, err := services.GptSimpleRequest(msg.Message)
	if err != nil {
		sendError(conn, "Error while calling MakeDecision service")
		return
	}

	res := websocketModels.MakeDecisionResponse{
		Message: back,
		Status:  200,
	}

	sendResponse(conn, res)
}
