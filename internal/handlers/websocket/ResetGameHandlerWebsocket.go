package websocketHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	websocketModels "my-api/internal/models/websocket"
	websocketServices "my-api/internal/services/websocket"
)

func ResetGameWebsocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var msg websocketModels.ResetGameRequest
	var initialRoute = "ResetGame"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	websocketServices.ResetGameServiceWebSocket(conn, msg, sendResponse, sendError)
}
