package apiTesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func TestWebSocketAndHTTPConnect() error {
	err := ResetTokenFile()

	isWSConnectSuccess := true
	isHttpConnectSuccess := true

	if err != nil {
		return err
	}

	wsToken, err := connectViaWebSocket()
	if err == nil {
		err = SaveToken("ws", wsToken)

		if err != nil {
			return err
		}
	} else {
		fmt.Printf("error in ws connect: %s", err)
		isWSConnectSuccess = false
	}

	httpToken, err := connectViaHTTP()
	if err == nil {
		err = SaveToken("http", httpToken)

		if err != nil {
			return err
		}
	} else {
		fmt.Printf("error in http connect: %s", err)
		isHttpConnectSuccess = false
	}

	if !isWSConnectSuccess && !isHttpConnectSuccess {
		return fmt.Errorf("connect failed")
	}
	return nil
}

func connectViaWebSocket() (string, error) {
	conn, _, err := websocket.DefaultDialer.Dial(WsConnectURL, nil)
	if err != nil {
		return "", fmt.Errorf("WebSocket dial error: %w", err)
	}
	defer conn.Close()

	message := map[string]string{
		"action":     "Connect",
		"identifier": WsID,
		"password":   Password,
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

	token := response["token"]
	if token == "" {
		return "", fmt.Errorf("missing token in response: %s", msg)
	}

	return token, nil
}

func connectViaHTTP() (string, error) {
	payload := map[string]string{
		"Identifier": HttpID,
		"Password":   Password,
	}

	body, _ := json.Marshal(payload)
	resp, err := http.Post(HttpBaseUrl+"Connect", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("JSON decode failed: %w", err)
	}

	token, ok := data["token"].(string)
	if !ok || token == "" {
		return "", fmt.Errorf("missing or invalid token: %v", data)
	}

	return token, nil
}
