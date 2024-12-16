package utils

import (
	"github.com/gorilla/websocket"
)

func SendResponse(conn *websocket.Conn, response string, initialRoute string) {
	resp := map[string]interface{}{
		"status": "success",
		"route":  initialRoute,
		"data":   response,
	}
	conn.WriteJSON(resp)
}

func SendError(conn *websocket.Conn, errorMessage string, initialRoute string) {
	resp := map[string]interface{}{
		"status":  "error",
		"route":   initialRoute,
		"message": errorMessage,
	}
	conn.WriteJSON(resp)
}
