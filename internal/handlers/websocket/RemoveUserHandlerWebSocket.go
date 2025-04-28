package websocketHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	"my-api/pkg"
)

// RemoveUserHandlerWebSocket handles WebSocket requests to remove a user based on a token and username.
func RemoveUserHandlerWebSocket(
	conn *websocket.Conn, message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var req sharedModel.RemoveUserRequest
	var initialRoute = "RemoveUser"

	err := json.Unmarshal(message, &req)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if req.Token == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the JSON body",
		})
		return
	}

	err = sharedServices.RemoveUserService(req.Token, req.UserName)
	if err != nil {
		pkg.DisplayContext("Internal Server Error", pkg.Error, err)
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Internal Server Error",
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"message": "Success",
	})
}
