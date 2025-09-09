package websocketServices

import (
	"strconv"

	"github.com/gorilla/websocket"

	"my-api/internal/services"
	"my-api/internal/types"
	"my-api/internal/utils"
	"my-api/pkg"
)

// RestartServiceWebSocket handles game restart for a specific user.
func RestartServiceWebSocket(
	conn *websocket.Conn,
	token string,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) {
	initialRoute := "Restart"

	userIDStr, err := utils.GetUserIDFromJWT(token)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{"message": "Invalid token"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		sendError(conn, initialRoute, "", map[string]interface{}{"message": "Invalid user id"})
		return
	}

	MarkUserResetting(userID)

	if err := services.DropEntitiesByUserID(userID); err != nil {
		UnmarkUserResetting(userID)
		sendError(conn, initialRoute, "", map[string]interface{}{"message": err.Error()})
		return
	}

	pkg.DeleteToken(userIDStr)

	sendResponse(conn, initialRoute, "", map[string]interface{}{"message": "Restarted"})

	conn.Close()
}