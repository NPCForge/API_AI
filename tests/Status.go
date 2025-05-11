package apiTesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type StatusResponse struct {
	Route   string `json:"route"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Status() error {
	err := StatusViaWebSocket()
	if err != nil {
		return err
	}

	err = StatusViaHTTP()
	if err != nil {
		return err
	}

	return nil
}

func StatusViaWebSocket() error {
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
		"action": "Status",
		"token":  token,
	}

	if err := conn.WriteJSON(message); err != nil {
		return fmt.Errorf("send JSON failed: %w", err)
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read message failed: %w", err)
	}

	var response StatusResponse
	if err := json.Unmarshal(msg, &response); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	status := response.Status

	if status != "success" {
		return fmt.Errorf("invalid status: %s", status)
	}

	return nil
}

func StatusViaHTTP() error {
	token, err := ReadTokens("http")

	if err != nil {
		return fmt.Errorf("http read token error: %w", err)
	}

	payload := map[string]string{}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("GET", HttpBaseUrl+"Status", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("JSON decode failed: %w", err)
	}

	status, ok := data["status"].(string)

	if !ok || status != "success" {
		return fmt.Errorf("invalid status: %s", status)
	}

	return nil
}
