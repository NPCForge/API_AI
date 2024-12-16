package websocketServices

import (
	"my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/pkg"
	"strconv"

	"github.com/gorilla/websocket"
)

func UserConnectWebSocket(conn *websocket.Conn, msg websocketModels.ConnectRequest, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	if msg.Action == "" || msg.Token == "" {
		sendError(conn, "Missing required fields in the JSON message")
		return
	}

	response, err := services.IsExist(msg.Token)

	if err != nil {
		sendError(conn, "Error searching in database")
		return
	}
	if !response {
		sendError(conn, "Account doesn't exist")
		return
	}

	id, err := services.GetIDFromDB(msg.Token)

	var stringId = strconv.Itoa(id)

	if err != nil {
		sendError(conn, "Error searching in database")
		return
	}

	pass, err := pkg.GenerateJWT(stringId)

	if err != nil {
		sendError(conn, "Error generating token")
		return
	}

	pkg.SetToken(stringId, pass)

	sendResponse(conn, pass)
}
