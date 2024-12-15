package services

import (
	"errors"
	"my-api/internal/models"
	"my-api/pkg"
	"strconv"

	"github.com/gorilla/websocket"
)

func SaveInDatabase(entity models.RegisterRequest) (int64, error) {
	response, err := IsExist(entity.Token)

	if err != nil {
		return -1, errors.New("error searching in table")
	}

	if response {
		return -1, errors.New("error entry already exist in database")
	}

	id, err := Register(entity.Token, entity)

	if err != nil {
		return -1, errors.New("error while registering")
	}

	return id, nil
}

func RegisterServiceWebSocket(conn *websocket.Conn, msg models.RegisterRequest, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
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

// Generation d'un token, et enregistrement de l'entité dans la base de donnée
func RegisterService(entity models.RegisterRequest) (string, error) {
	id, err := SaveInDatabase(entity)

	if err != nil {
		return "", errors.New("error saving in database")
	}

	pass, err := pkg.GenerateJWT(strconv.FormatInt(id, 10))

	if err != nil {
		return "", errors.New("error generating JWT")
	}

	return pass, nil
}
