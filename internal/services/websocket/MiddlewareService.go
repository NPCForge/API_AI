package websocketServices

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"my-api/internal/utils"
	"my-api/pkg"
)

func LoginMiddlewareWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) bool {
	var msg struct {
		Action string `json:"action"`
		Token  string `json:"token"`
	}
	var initialRoute = "Connection"

	err := json.Unmarshal(message, &msg)
	if err != nil {
		utils.SendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return false
	}

	if msg.Token == "" {
		utils.SendError(conn, initialRoute, map[string]interface{}{
			"message": "No token in request body",
		})
		return false
	}

	_, err = pkg.VerifyJWT(msg.Token)

	if err != nil {
		utils.SendError(conn, initialRoute, map[string]interface{}{
			"message": "Invalid Token",
		})
		return false
	}

	return true
}
