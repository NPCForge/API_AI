package apiTesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func MakeDecision() error {

	fmt.Println("MakeDecision WS")
	err := makeDecisionViaWebSocket()
	if err != nil {
		fmt.Println("Error removing user via WebSocket:", err)
		return err
	}

	fmt.Println("MakeDecision HTTP")
	err = makeDecisionViaHTTP()
	if err != nil {
		fmt.Println("Error removing user via Http:", err)
		return err
	}
	return nil
}

func makeDecisionViaWebSocket() error {
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
		"action":   "MakeDecision",
		"checksum": "WsChecksum",
		"message":  "Nearby Entities: {[Checksum = WsChecksum]}",
		"token":    token,
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
		return fmt.Errorf("received error: %s", msg)
	}

	return nil
}

func makeDecisionViaHTTP() error {
	token, err := ReadTokens("http")

	if err != nil {
		return fmt.Errorf("http read token error: %w", err)
	}

	payload := map[string]string{
		"checksum": "HttpChecksum",
		"message":  "Nearby Entities: {[Checksum = HttpChecksum]}",
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", HttpBaseUrl+"MakeDecision", bytes.NewBuffer(body))
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

	status, ok := data["status"].(float64)
	if !ok || status != 200 {
		return fmt.Errorf("invalid status received: %v", data)
	}

	return nil
}
