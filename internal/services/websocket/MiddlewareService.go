package websocketServices

import (
	"encoding/json"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"

	"my-api/internal/types"
	"my-api/internal/utils"
)

// LoginMiddlewareWebSocket verifies the JWT token in an incoming WebSocket message before proceeding with protected actions.
func LoginMiddlewareWebSocket(
	conn *websocket.Conn,
	message []byte,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) bool {
	var msg struct {
		Action string `json:"action"`
		Token  string `json:"token"`
	}

	initialRoute := "Connection"

	color.Cyan("🔐 WebSocket login middleware triggered")

	err := json.Unmarshal(message, &msg)
	if err != nil {
		color.Red("❌ Failed to decode JSON: %v", err)
		utils.SendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Error while decoding JSON message",
		})
		return false
	}

	if msg.Token == "" {
		color.Yellow("⚠️ Token missing in request body")
		utils.SendError(conn, initialRoute, "", map[string]interface{}{
			"message": "No token provided in request body",
		})
		return false
	}

	color.Yellow("🔐 Token: " + msg.Token)

	_, err = utils.VerifyJWT(msg.Token)
	if err != nil {
		color.Red("❌ Invalid JWT: %v", err)
		utils.SendError(conn, initialRoute, "", map[string]interface{}{
			"message": "Invalid token",
		})
		return false
	}

	color.Green("✅ JWT verified successfully")
	return true
}
