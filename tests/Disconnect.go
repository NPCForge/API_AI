package apiTesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func Disconnect() error {
	err := disconnectViaWebSocket()
	if err != nil {
		return err
	}

	err = disconnectEntityViaHTTP()
	if err != nil {
		return err
	}

	return nil
}

func disconnectViaWebSocket() error {
	conn, _, err := websocket.DefaultDialer.Dial(WsConnectURL, nil)
	if err != nil {
		return fmt.Errorf("WebSocket dial error: %w", err)
	}
	defer conn.Close()

	token, err := ReadTokens("ws")

	if err != nil {
		return fmt.Errorf("WebSocket token read error: %w", err)
	}

	message := map[string]string{
		"action": "Disconnect",
		"token":  token,
	}

	if err := conn.WriteJSON(message); err != nil {
		return fmt.Errorf("send JSON failed: %w", err)
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read message failed: %w", err)
	}

	var response map[string]string
	if err := json.Unmarshal(msg, &response); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	status := response["status"]
	if status != "success" {
		return fmt.Errorf("missing success in response: %s", msg)
	}

	return nil
}

func disconnectEntityViaHTTP() error {
	token, err := ReadTokens("http")

	if err != nil {
		return fmt.Errorf("http read token error: %w", err)
	}

	payload := map[string]string{}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", HttpBaseUrl+"Disconnect", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token) // ✅ correct

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
