package websocketHandlers

import (
	"encoding/json"
	"my-api/internal/models/websocket"
	"my-api/internal/services/websocket"

	"github.com/gorilla/websocket"
)

func RegisterHandlerWebsocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg websocketModels.RegisterRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

	if msg.Action == "" || msg.Token == "" || msg.Name == "" || msg.Prompt == "" {
		sendError(conn, "Missing required fields in the JSON message")
		return
	}

	websocketServices.RegisterServiceWebSocket(conn, msg, sendResponse, sendError)
}
