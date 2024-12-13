package services

import (
	"encoding/json"
	"errors"

	"my-api/internal/models"

	"github.com/gorilla/websocket"
)

func MakeDecisionWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg models.MakeDecisionRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

	if msg.Message == "" {
		sendError(conn, "Missing required fields in the JSON message")
		return
	}

	back, err := MakeDecisionService(msg.Message)
	if err != nil {
		sendError(conn, "Error while calling MakeDecision service")
		return
	}

	res := models.MakeDecisionResponse{
		Message: back,
		Status:  200,
	}

	sendResponse(conn, res)
}

func MakeDecisionService(req string) (string, error) {
	res, err := GptSimpleRequest(req)

	if err != nil {
		return "", errors.New("error requesting Chatgpt")
	}

	return res, nil
}
