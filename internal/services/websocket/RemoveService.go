package websocketServices

import (
	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/pkg"

	"github.com/gorilla/websocket"
)

func RemoveServiceWebSocket(conn *websocket.Conn, msg websocketModels.RemoveRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "Remove"

	UserId, err := pkg.GetUserIDFromJWT(msg.Token)

	if err != nil {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "error during the process",
		})
	}

	exist, err_ := services.IsExistById(UserId)

	if err_ != nil {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "failed",
		})
		return
	}

	if !exist {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "success",
		})
		return
	}

	response, err_ := services.DropUser(UserId)

	if err_ != nil || response == "" {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "failed: droping BD",
		})
		return
	}

	sendResponse(conn, initialRoute, map[string]interface{}{
		"message": "success",
	})
}
