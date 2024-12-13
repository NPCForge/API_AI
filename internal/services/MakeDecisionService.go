package services

import (
    "encoding/json"
	"log"
	"errors"

    "my-api/internal/models"
    
	"github.com/gorilla/websocket"
)

func MakeDecisionWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	// Create request from message
    var msg struct {
        Action  string `json:"action"`
        Message string `json:"message"`
    }	
    
    err := json.Unmarshal(message, &msg)
    if err != nil {
        sendError(conn, "Error while decoding JSON message")
        return
    }
    	
	req := models.MakeDecisionRequest{
		Message: msg.Message,
	}

	// Call MakeDecision service
	back, err := MakeDecisionService(req.Message)
	if err != nil {
		sendError(conn, "Error while calling MakeDecision service")
		return
	}

	// Create response
	res := models.MakeDecisionResponse{
		Message: back,
		Status:  200,
	}

	log.Printf("back = %s", back)

	// Send websocket response
	sendResponse(conn, res)
}

func MakeDecisionService(req string) (string, error) {
	res, err := GptSimpleRequest(req)

	if err != nil {
		return "", errors.New("error requesting Chatgpt")
	}

	return res, nil
}
