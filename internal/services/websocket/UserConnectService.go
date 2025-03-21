package websocketServices

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"

	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/internal/types"
	"my-api/pkg"
)

func UserConnectWebSocket(
	conn *websocket.Conn,
	msg websocketModels.ConnectRequest,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) {
	initialRoute := "Connection"

	color.Cyan("🔌 UserConnectWebSocket triggered")

	if msg.Action == "" || msg.Token == "" {
		color.Yellow("⚠️ Missing required fields in request")
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Missing required fields in the JSON message",
		})
		return
	}

	exists, err := services.IsExist(msg.Token)
	if err != nil {
		color.Red("❌ Error checking existence in DB: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error searching in database",
		})
		return
	}
	if !exists {
		color.Yellow("🚫 Account does not exist for token: %s", msg.Token)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Account doesn't exist",
		})
		return
	}

	id, err := services.GetIDFromDB(msg.Token)
	if err != nil {
		color.Red("❌ Failed to get ID from DB: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error searching in database",
		})
		return
	}
	stringId := strconv.Itoa(id)

	pass, err := pkg.GenerateJWT(stringId)
	if err != nil {
		color.Red("❌ Failed to generate JWT: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error generating token",
		})
		return
	}

	pkg.SetToken(stringId, pass)
	color.Green("✅ User %s connected and token generated", stringId)

	sendResponse(conn, initialRoute, map[string]interface{}{
		"token": pass,
	})
}
