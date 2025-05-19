package apiTesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func CreateEntity() error {
	wsID, err := createEntityViaWebSocket()
	if err != nil {
		return err
	}

	fmt.Printf("created entity wsID: %s\n", wsID)

	httpID, err := createEntityViaHTTP()
	if err != nil {
		return err
	}

	fmt.Printf("created entity httpID: %s\n", httpID)
	return nil
}

func createEntityViaWebSocket() (string, error) {
	conn, _, err := websocket.DefaultDialer.Dial(WsConnectURL, nil)
	if err != nil {
		return "", fmt.Errorf("WebSocket dial error: %w", err)
	}
	defer conn.Close()

	token, err := ReadTokens("ws")

	if err != nil {
		return "", fmt.Errorf("WebSocket token read error: %w", err)
	}

	message := map[string]string{
		"action":   "CreateEntity",
		"name":     "WsEntity",
		"prompt":   "WsPrompt",
		"checksum": "WsChecksum",
		"role":     "WsRole",
		"token":    token,
	}

	if err := conn.WriteJSON(message); err != nil {
		return "", fmt.Errorf("send JSON failed: %w", err)
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		return "", fmt.Errorf("read message failed: %w", err)
	}

	var response map[string]string
	if err := json.Unmarshal(msg, &response); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	id := response["id"]
	if id == "" {
		return "", fmt.Errorf("missing id in response: %s", msg)
	}

	return id, nil
}

func createEntityViaHTTP() (string, error) {
	token, err := ReadTokens("http")

	if err != nil {
		return "", fmt.Errorf("http read token error: %w", err)
	}

	payload := map[string]string{
		"name":     "HttpEntity",
		"prompt":   "HttpPrompt",
		"role":     "HttpRole",
		"checksum": "HttpChecksum",
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", HttpBaseUrl+"CreateEntity", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token) // ✅ correct

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("JSON decode failed: %w", err)
	}

	httpId, ok := data["id"].(string)
	if !ok || httpId == "" {
		return "", fmt.Errorf("missing or invalid id: %v", data)
	}

	return httpId, nil
}
