package websocketHandlers

import (
	"encoding/json"
	sharedModel "my-api/internal/models/shared"
	service "my-api/internal/services/merged"
	"my-api/pkg"

	"github.com/gorilla/websocket"
)

func RemoveHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, map[string]interface{}),
) {
	var req sharedModel.RemoveUserRequest
	var initialRoute = "Remove"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if req.DeleteUserIdentifier == "" || req.Token == "" {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Bad Request",
		})
		return
	}

	disconnect, err := service.RemoveService(req.Token, req.DeleteUserIdentifier)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Internal Server Error",
		})
		return
	}

	if disconnect {
		pkg.DisplayContext("Remove Handler need to Disconnect WS", pkg.Debug)
	}

	sendResponse(conn, initialRoute, map[string]interface{}{
		"message": "Success",
	})
}
