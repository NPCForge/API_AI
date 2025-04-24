package exemples

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

func NameHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, map[string]interface{}),
) {
	var req Request
	var initialRoute = "Name"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	// pass, err := NameService()

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	sendResponse(conn, initialRoute, map[string]interface{}{
		"token": "",
	})
}
