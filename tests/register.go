package apiTesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func TestWebsocketAndHTTPRegister() error {
	err := ResetTokenFile()

	isWSRegisterSuccess := true
	isHTTPRegisterSuccess := true

	if err != nil {
		return err
	}

	// WebSocket test
	wsToken, err := testWebSocketRegister()
	if err == nil {
		err = SaveToken("ws", wsToken)

		if err != nil {
			return err
		}
	} else {
		fmt.Println("Error in websocket register")
		isWSRegisterSuccess = false
	}

	// HTTP test
	httpToken, err := testHTTPRegister()
	if err == nil {
		err = SaveToken("http", httpToken)

		if err != nil {
			return err
		}
	} else {
		fmt.Println("Error in http register")
		isHTTPRegisterSuccess = false
	}

	if !isWSRegisterSuccess && isHTTPRegisterSuccess {
		return fmt.Errorf("all connexions failed")
	}

	return nil
}

func testWebSocketRegister() (string, error) {
	conn, _, err := websocket.DefaultDialer.Dial(WsConnectURL, nil)
	if err != nil {
		return "", fmt.Errorf("dial error: %w", err)
	}
	defer conn.Close()

	msg := map[string]string{
		"action":     "Register",
		"API_KEY":    ApiKey,
		"identifier": WsID,
		"password":   Password,
	}

	if err := conn.WriteJSON(msg); err != nil {
		return "", fmt.Errorf("send error: %w", err)
	}

	_, data, err := conn.ReadMessage()
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}

	var resp map[string]string
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", fmt.Errorf("unmarshal error: %w", err)
	}

	token := resp["token"]
	if token == "" {
		return "", fmt.Errorf("no token in response: %v", string(data))
	}
	return token, nil
}

func testHTTPRegister() (string, error) {
	reqData := map[string]string{
		"API_KEY":    ApiKey,
		"Identifier": HttpID,
		"Password":   Password,
	}
	payload, _ := json.Marshal(reqData)

	resp, err := http.Post(HttpBaseUrl+"Register", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("post error: %w", err)
	}
	defer resp.Body.Close()

	var respData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return "", fmt.Errorf("decode error: %w", err)
	}

	token, ok := respData["token"].(string)
	if !ok || token == "" {
		return "", fmt.Errorf("no token in HTTP response: %v", respData)
	}

	return token, nil
}
