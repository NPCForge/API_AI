package websocketServices

import (
	"github.com/fatih/color"
	"github.com/gorilla/websocket"

	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/internal/types"
	"my-api/pkg"
)

func RemoveServiceWebSocket(
	conn *websocket.Conn,
	msg websocketModels.RemoveRequest,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) {
	initialRoute := "Remove"

	color.Cyan("🗑️  RemoveServiceWebSocket triggered")

	UserId, err := pkg.GetUserIDFromJWT(msg.Token)
	if err != nil {
		color.Red("❌ Failed to extract user ID from token: %v", err)
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "error during the process",
		})
		return
	}

	exist, err_ := services.IsExistById(UserId)
	if err_ != nil {
		color.Red("❌ Failed to check if user exists: %v", err_)
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "failed",
		})
		return
	}

	if !exist {
		color.Yellow("⚠️ User ID %d does not exist", UserId)
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "success",
		})
		return
	}

	response, err_ := services.DropUser(UserId)
	if err_ != nil || response == "" {
		color.Red("❌ Failed to drop user ID %d: %v", UserId, err_)
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "failed: dropping DB",
		})
		return
	}

	color.Green("✅ User ID %d successfully removed", UserId)
	sendResponse(conn, initialRoute, map[string]interface{}{
		"message": "success",
	})
}
