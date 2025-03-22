package websocketServices

import (
	"github.com/fatih/color"
	"github.com/gorilla/websocket"

	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/internal/types"
	"my-api/pkg"
)

func NewMessageWebSocket(
	conn *websocket.Conn,
	msg websocketModels.NewMessageRequest,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) {
	initialRoute := "NewMessage"

	color.Cyan("✉️  Handling new message from: %s", msg.Sender)

	receiverId, err := pkg.GetUserIDFromJWT(msg.Token)
	if err != nil {
		color.Red("❌ Failed to extract receiver ID from JWT: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Invalid token",
		})
		return
	}

	senderId, err := services.GetIDFromDB(msg.Sender)
	if err != nil {
		color.Red("❌ Failed to get sender ID: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error in getting IDs",
		})
		return
	}

	_, err = services.NewMessage(senderId, receiverId, msg.Message)
	if err != nil {
		color.Red("❌ Error while creating new message: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error in creating message",
		})
		return
	}

	color.Green("✅ Message successfully saved from %d to %d", senderId, receiverId)

	sendResponse(conn, initialRoute, map[string]interface{}{
		"message": "Message sent successfully",
	})
}
