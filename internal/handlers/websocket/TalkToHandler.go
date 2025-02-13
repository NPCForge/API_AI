package websocketHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	websocketModels "my-api/internal/models/websocket"
	websocketServices "my-api/internal/services/websocket"
)

func TalkToHandlerWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var msg websocketModels.TalkToRequest
	var initialRoute = "TalkTo"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if msg.Message == "" {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	websocketServices.TalkToWebSocket(conn, msg, sendResponse, sendError)
}
