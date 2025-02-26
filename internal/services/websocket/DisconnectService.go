package websocketServices

import (
	"github.com/gorilla/websocket"
	"my-api/internal/models/websocket"
	"my-api/pkg"
)

func DisconnectWebSocket(conn *websocket.Conn, msg websocketModels.DisconnectRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "Disconnect"
	result, err := pkg.VerifyJWT(msg.Token)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while verifying JWT",
		})
		return
	}

	if result != nil {
		pkg.DeleteToken(result.UserID)
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "Disconnected",
		})
	} else {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Failed to disconnect",
		})
	}

}
