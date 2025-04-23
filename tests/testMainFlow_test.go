package apiTesting

import (
	"fmt"
	"testing"
)

func TestMainFlow(t *testing.T) {
	fmt.Println("Testing register...")

	var isConnexionSuccess = false
	var err error

	t.Run("Step1_Register", func(t *testing.T) {
		if err := TestWebsocketAndHTTPRegister(); err == nil {
			isConnexionSuccess = true
		}
	})

	fmt.Println("Testing connect...")

	t.Run("Step2_Connect", func(t *testing.T) {
		if err := TestWebSocketAndHTTPConnect(); err == nil {
			isConnexionSuccess = true
		}
	})

	if !isConnexionSuccess {
		t.Fatalf("Connect failed: %v", err)
	}

}
