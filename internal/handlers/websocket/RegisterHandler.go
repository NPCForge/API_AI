package websocketHandlers

import (
	"encoding/json"
	"my-api/config"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

func RegisterHandlerWebsocket(
	conn *websocket.Conn,
	message []byte,
	sendResponse func(*websocket.Conn, string, string, map[string]interface{}),
	sendError func(*websocket.Conn, string, string, map[string]interface{}),
) {
	var msg sharedModel.RegisterRequest
	var initialRoute = "Register"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return
	}

	if msg.Action == "" || msg.Token == "" || msg.Identifier == "" || msg.Password == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	apiKey := config.GetEnvVariable("API_KEY_REGISTER")

	if msg.Token != apiKey {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Invalid API Key",
		})
		return
	}

	private, id, err := sharedServices.RegisterService(msg.Password, msg.Identifier)

	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	sendResponse(conn, initialRoute, "", map[string]interface{}{
		"token": private,
		"id":    id,
	})
}
