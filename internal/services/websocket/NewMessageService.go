package websocketServices

import (
	"github.com/gorilla/websocket"
	"my-api/internal/models/websocket"
	"my-api/internal/services"
)

func NewMessageWebSocket(conn *websocket.Conn, msg websocketModels.NewMessageRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "NewMessage"

	senderId, err := services.GetIDFromDB(msg.Sender)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error in getting IDs",
		})
		return
	}

	for _, receiverChecksum := range msg.Receivers {
		receiverId, err := services.GetIDFromDB(receiverChecksum)
		if err != nil {
			sendError(conn, initialRoute, map[string]interface{}{
				"message": "Error in getting IDs",
			})
			return
		}

		_, err = services.NewMessage(senderId, receiverId, msg.Message)

		if err != nil {
			sendError(conn, initialRoute, map[string]interface{}{
				"message": "Error in creating message",
			})
		}
	}

	return
}
