package websocketServices

import (
	"my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/pkg"
	"strconv"

	"github.com/gorilla/websocket"
)

func UserConnectWebSocket(conn *websocket.Conn, msg websocketModels.ConnectRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "Connection"

	if msg.Action == "" || msg.Token == "" {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	response, err := services.IsExist(msg.Token)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error searching in database",
		})
		return
	}
	if !response {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Account doesn't exist",
		})
		return
	}

	id, err := services.GetIDFromDB(msg.Token)

	var stringId = strconv.Itoa(id)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error searching in database",
		})
		return
	}

	pass, err := pkg.GenerateJWT(stringId)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error generating token",
		})
		return
	}

	pkg.SetToken(stringId, pass)

	sendResponse(conn, initialRoute, map[string]interface{}{
		"token": pass,
	})
}
