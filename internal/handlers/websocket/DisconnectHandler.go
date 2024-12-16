package websocketHandlers

import (
	"encoding/json"
	"my-api/internal/models/websocket"
	"my-api/internal/services/websocket"

	"github.com/gorilla/websocket"
)

func DisconnectHandlerWebsocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, string, string), sendError func(*websocket.Conn, string, string)) {
	var msg websocketModels.DisconnectRequest
	var initialRoute = "Disconnect"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message", initialRoute)
		return
	}

	websocketServices.DisconnectWebSocket(conn, msg, sendResponse, sendError)
}
