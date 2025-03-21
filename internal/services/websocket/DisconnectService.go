package websocketServices

import (
	"github.com/fatih/color"
	"github.com/gorilla/websocket"

	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/types"
	"my-api/pkg"
)

func DisconnectWebSocket(
	conn *websocket.Conn,
	msg websocketModels.DisconnectRequest,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) {
	initialRoute := "Disconnect"

	color.Cyan("🔌 Attempting to disconnect...")

	result, err := pkg.VerifyJWT(msg.Token)
	if err != nil {
		color.Red("❌ JWT verification failed: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while verifying JWT",
		})
		return
	}

	if result != nil {
		color.Green("✅ Valid token. Disconnecting user ID: %v", result.UserID)
		pkg.DeleteToken(result.UserID)
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "Disconnected",
		})
	} else {
		color.Yellow("⚠️ Token is nil after verification")
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Failed to disconnect",
		})
	}
}
