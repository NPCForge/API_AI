package websocketServices

import (
	"github.com/fatih/color"
	"github.com/gorilla/websocket"

	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/internal/types"
	"strconv"
)

func NewMessageWebSocket(
	conn *websocket.Conn,
	msg websocketModels.NewMessageRequest,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) {
	initialRoute := "NewMessage"

	color.Cyan("✉️  Handling new message from: %s", msg.Sender)

	senderId, err := services.GetIDFromDB(msg.Sender)
	if err != nil {
		color.Red("❌ Failed to get sender ID: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error in getting IDs",
		})
		return
	}

	for _, receiver := range msg.Receivers {
		receiverId, err := services.GetEntityByName(receiver)
		if err != nil {
			color.Red("❌ Error while getting IDs: %v", err)
			sendError(conn, initialRoute, map[string]interface{}{
				"message": "Error in getting IDs",
			})
			return
		}

		receiverIntId, err := strconv.Atoi(receiverId)

		if err != nil {
			color.Red("❌ Error while creating new message: %v", err)
			sendError(conn, initialRoute, map[string]interface{}{
				"message": "Error in creating message",
			})
			return
		}

		_, err = services.NewMessage(senderId, receiverIntId, msg.Message)

		if err != nil {
			sendError(conn, initialRoute, map[string]interface{}{
				"message": "Error in creating message",
			})
			return
		}
		color.Green("✅ Message successfully saved from %d to %d", senderId, receiverId)
	}

	return
}
