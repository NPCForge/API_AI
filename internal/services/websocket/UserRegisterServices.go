package websocketServices

import (
	"errors"
	"my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/pkg"
	"strconv"

	"github.com/gorilla/websocket"
)

func SaveInDatabase(entity websocketModels.RegisterRequest) (int64, error) {
	response, err := services.IsExist(entity.Token)

	if err != nil {
		return -1, errors.New("error searching in table")
	}

	if response {
		return -1, errors.New("error entry already exist in database")
	}

	id, err := services.RegisterWebsocket(entity.Token, entity)

	if err != nil {
		return -1, errors.New("error while registering")
	}

	return id, nil
}

func RegisterServiceWebSocket(conn *websocket.Conn, msg websocketModels.RegisterRequest, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	id, err := SaveInDatabase(msg)

	var stringId = strconv.FormatInt(id, 10)

	if err != nil {
		sendError(conn, "Error saving in database")
		return
	}

	pass, err := pkg.GenerateJWT(stringId)

	if err != nil {
		sendError(conn, "Unable to generate Token")
		return
	}

	pkg.SetToken(stringId, pass)

	sendResponse(conn, pass)
}
