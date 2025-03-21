package utils

import (
	"github.com/gorilla/websocket"
)

func SendResponse(conn *websocket.Conn, initialRoute string, fields map[string]interface{}) {
	resp := map[string]interface{}{
		"status": "success",
		"route":  initialRoute,
	}

	for key, value := range fields {
		resp[key] = value
	}

	println("Sending success response:")
	println(resp)

	conn.WriteJSON(resp)
}

func SendError(conn *websocket.Conn, initialRoute string, fields map[string]interface{}) {
	resp := map[string]interface{}{
		"status": "error",
		"route":  initialRoute,
	}

	for key, value := range fields {
		resp[key] = value
	}

	println("Sending error response:")
	println(resp)

	conn.WriteJSON(resp)
}
