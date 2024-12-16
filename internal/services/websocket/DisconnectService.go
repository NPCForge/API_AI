package websocketServices

import (
	"github.com/gorilla/websocket"
	"my-api/internal/models/websocket"
	"my-api/pkg"
)

func DisconnectWebSocket(conn *websocket.Conn, msg websocketModels.DisconnectRequest, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	result, err := pkg.VerifyJWT(msg.Token)

	if err != nil {
		sendError(conn, "Error while verifying JWT")
		return
	}

	if result != nil {
		pkg.DeleteToken(result.UserID)
		sendResponse(conn, "Disconnected")
	} else {
		sendError(conn, "Failed to disconnect")
	}

}
