package websocket

import (
	"encoding/json"
	"my-api/internal/models"
	"my-api/internal/services"

	"github.com/gorilla/websocket"
)

func RegisterHandlerWebsocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg models.RegisterRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

	if msg.Action == "" || msg.Token == "" || msg.Name == "" || msg.Prompt == "" {
		sendError(conn, "Missing required fields in the JSON message")
		return
	}

	services.RegisterServiceWebSocket(conn, msg, sendResponse, sendError)
}
