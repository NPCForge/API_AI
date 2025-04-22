package websocketHandlers

import (
	"encoding/json"
	websocketModels "my-api/internal/models/websocket"
	websocketServices "my-api/internal/services/websocket"

	"github.com/gorilla/websocket"
)

func RemoveHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, map[string]interface{}),
) {
	var msg websocketModels.RemoveRequest
	var initialRoute = "Remove"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	websocketServices.RemoveServiceWebSocket(conn, msg, sendResponse, sendError)
}
