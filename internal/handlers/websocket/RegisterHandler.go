package websocketHandlers

import (
	"encoding/json"
	"my-api/config"
	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"

	"github.com/gorilla/websocket"
)

// RegisterHandlerWebsocket handles WebSocket requests to register a new user after API key verification.
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

	if msg.Action == "" || msg.Token == "" || msg.Identifier == "" || msg.Password == "" || msg.GamePrompt == "" {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Missing required fields in the JSON body",
		})
		return
	}

	apiKey := config.GetEnvVariable("API_KEY_REGISTER")

	if msg.Token != apiKey {
		sendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Invalid API key",
		})
		return
	}

	private, id, err := sharedServices.RegisterService(msg.Password, msg.Identifier, msg.GamePrompt)
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
