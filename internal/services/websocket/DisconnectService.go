package websocketServices

import (
	"github.com/gorilla/websocket"
	"my-api/internal/models/websocket"
	"my-api/pkg"
)

func DisconnectWebSocket(conn *websocket.Conn, msg websocketModels.DisconnectRequest, sendResponse func(*websocket.Conn, string, string), sendError func(*websocket.Conn, string, string)) {
	var initialRoute = "Disconnect"
	result, err := pkg.VerifyJWT(msg.Token)

	if err != nil {
		sendError(conn, "Error while verifying JWT", initialRoute)
		return
	}

	if result != nil {
		pkg.DeleteToken(result.UserID)
		sendResponse(conn, "Disconnected", initialRoute)
	} else {
		sendError(conn, "Failed to disconnect", initialRoute)
	}

}
