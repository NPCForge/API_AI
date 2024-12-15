package services

import (
	"encoding/json"
	"errors"
	"my-api/internal/models"
	"my-api/pkg"
	"strconv"

	"github.com/gorilla/websocket"
)

func UserConnectWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg models.ConnectRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

	if msg.Action == "" || msg.Token == "" {
		sendError(conn, "Missing required fields in the JSON message")
		return
	}

	response, err := IsExist(msg.Token)

	if err != nil {
		sendError(conn, "Error searching in database")
		return
	}
	if !response {
		sendError(conn, "Account doesn't exist")
		return
	}

	id, err := GetIDFromDB(msg.Token)

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

func UserConnect(req models.ConnectRequest) (string, error) {
	response, err := IsExist(req.Token)

	if err != nil {
		return "", errors.New("error searching in table")
	}

	if !response {
		return "", errors.New("account didn't exist")
	}

	pass, err := pkg.GenerateJWT(req.Token)

	pkg.SetToken(req.Token, pass)

	if err != nil {
		return "", errors.New("error generating JWT")
	}

	return pass, nil
}
