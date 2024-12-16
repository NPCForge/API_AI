package websocketServices

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"my-api/pkg"
)

func LoginMiddlewareWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) bool {
	var msg struct {
		Action string `json:"action"`
		Token  string `json:"token"`
	}

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return false
	}

	if msg.Token == "" {
		sendError(conn, "No token in request body")
		return false
	}

	_, err = pkg.VerifyJWT(msg.Token)

	if err != nil {
		sendError(conn, "Invalid Token")
		return false
	}

	return true
}
