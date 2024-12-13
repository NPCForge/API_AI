package services

import (
	"encoding/json"
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

func RegisterServiceWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg models.RegisterRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

    if msg.Action == "" || msg.Token == "" || msg.Name == "" || msg.Prompt == "" {
        sendError(conn, "Missing required fields in the JSON message")
        return
    }

	id, err := SaveInDatabase(msg)

	if err != nil {
		sendError(conn, "Error saving in database")
		return
	}

	pass, err := pkg.GenerateJWT(strconv.FormatInt(id, 10))

	if err != nil {
		sendError(conn, "Unable to generate Token")
		return
	}

	sendResponse(conn, pass)
}

// Generation d'un token, et enregistrement de l'entité dans la base de donnée
// func RegisterService(entity models.RegisterRequest) (string, error) {
// 	pass, err := pkg.GenerateJWT(entity.Name)

// 	if err != nil {
// 		return "", errors.New("error generating JWT")
// 	}

// 	pass, err = SaveInDatabase(pass, entity)

// 	if err != nil {
// 		return "", errors.New("error saving in database")
// 	}

// 	return pass, nil
// }
