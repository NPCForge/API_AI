package websocketHandlers

import (
	"encoding/json"
	"my-api/internal/models/websocket"
	"my-api/internal/services/websocket"

	"github.com/gorilla/websocket"
)

func ConnectHandlerWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, string, string), sendError func(*websocket.Conn, string, string)) {
	var msg websocketModels.ConnectRequest
	var initialRoute = "Connection"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message", initialRoute)
		return
	}

	websocketServices.UserConnectWebSocket(conn, msg, sendResponse, sendError)
}
