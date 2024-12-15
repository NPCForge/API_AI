package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"my-api/internal/models"
	"my-api/pkg"
)

func DisconnectWebSocket(conn *websocket.Conn, message []byte, sendResponse func(*websocket.Conn, interface{}), sendError func(*websocket.Conn, string)) {
	var msg models.DisconnectRequest

	err := json.Unmarshal(message, &msg)
	if err != nil {
		sendError(conn, "Error while decoding JSON message")
		return
	}

	if msg.Token == "" {
		sendError(conn, "Missing required fields in the JSON message")
		return
	}

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

func Disconnect(token string) (string, error) {
	userid, err := pkg.GetUserID(token)

	if !err {
		return "failed", errors.New("error getting userid")
	}

	res := pkg.DeleteToken(userid)

	if !res {
		fmt.Printf("Disconnect")
		return "failed", errors.New("error didn't exist")
	}

	return "success", nil
}
