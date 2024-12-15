package websocket

import (
	"encoding/json"
	"my-api/internal/models"
	"my-api/internal/services"

	"github.com/gorilla/websocket"
)

func DisconnectHandlerWebsocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg models.DisconnectRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

	services.DisconnectWebSocket(conn, msg, sendResponse, sendError)
}
