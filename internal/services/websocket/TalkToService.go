package websocketServices

import (
	"my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/pkg"

	"github.com/gorilla/websocket"
)

func TalkToWebSocket(conn *websocket.Conn, msg websocketModels.TalkToRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "TalkTo"

	UserId, err := pkg.GetUserIDFromJWT(msg.Token)

	if err != nil {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "error during the process",
		})
	}

	_, err_ := services.IsExistById(UserId)

	if err_ != nil {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "failed",
		})
		return
	}

	response, err_ := services.GetPromptByID(UserId)

	if err_ != nil {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "failed",
		})
		return
	}

	back, err := services.GptTalkToRequest(msg.Message, response)
	if err != nil {
		println("Error in TalkToWebSocket: " + err.Error())
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while calling TalkTo service",
		})
		return
	}

	sendResponse(conn, initialRoute, map[string]interface{}{
		"message": back,
	})
}
