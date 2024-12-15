package websocket

import (
	"encoding/json"
	"my-api/internal/models"
	"my-api/internal/services"

	"github.com/gorilla/websocket"
)

func ConnectHandlerWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg models.ConnectRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

	services.UserConnectWebSocket(conn, msg, sendResponse, sendError)
}
