package websocketHandlers

import (
	"github.com/gorilla/websocket"
)

func StatusHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var initialRoute = "Status"

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"message": "Authentified",
		"status":  200,
	})
}
