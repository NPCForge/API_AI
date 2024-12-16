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
	response, err := services.IsExist(entity.Checksum)

	if err != nil {
		return -1, errors.New("error searching in table")
	}

	if response {
		return -1, errors.New("error entry already exist in database")
	}

	id, err := services.RegisterWebsocket(entity.Checksum, entity)

	if err != nil {
		return -1, errors.New("error while registering")
	}

	return id, nil
}

func RegisterServiceWebSocket(conn *websocket.Conn, msg websocketModels.RegisterRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "Register"
	id, err := SaveInDatabase(msg)

	var stringId = strconv.FormatInt(id, 10)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error saving in database",
		})
		return
	}

	pass, err := pkg.GenerateJWT(stringId)

	if err != nil {
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Unable to generate Token",
		})
		return
	}

	pkg.SetToken(stringId, pass)

	sendResponse(conn, initialRoute, map[string]interface{}{
		"token": pass,
	})
}
