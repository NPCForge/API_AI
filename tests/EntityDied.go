package apiTesting

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func EntityDied() error {
	return entityDiedViaWebSocket()
	// No HTTP test for now as the user issue was specifically about "Cannot find matching route" which implies WS in this context (as HTTP has explicit route)
	// But actually the error message "Cannot find matching route for: EntityDied" is printed by the WebSocket handler's default case.
}

func entityDiedViaWebSocket() error {
	conn, _, err := websocket.DefaultDialer.Dial(WsConnectURL, nil)
	if err != nil {
		return fmt.Errorf("WebSocket dial error: %w", err)
	}
	defer conn.Close()

	token, err := ReadTokens("ws")
	if err != nil {
		return fmt.Errorf("WebSocket token read error: %w", err)
	}

	// We'll try to kill the entity created by CreateEntity.go ("WsChecksum")
	message := map[string]string{
		"action":   "EntityDied",
		"checksum": "WsChecksum",
		"token":    token,
	}

	if err := conn.WriteJSON(message); err != nil {
		return fmt.Errorf("send JSON failed: %w", err)
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read message failed: %w", err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(msg, &response); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	status, ok := response["status"].(float64) // JSON numbers are floats
	if !ok || status != 200 {
		return fmt.Errorf("unexpected status or response: %v", response)
	}

	messageStr, _ := response["message"].(string)
	if messageStr != "Entity death processed successfully" {
		return fmt.Errorf("unexpected success message: %s", messageStr)
	}

	return nil
}
