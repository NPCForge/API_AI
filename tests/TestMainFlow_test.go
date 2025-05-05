package apiTesting

import (
	"fmt"
	"testing"
)

func TestMainFlow(t *testing.T) {
	fmt.Println("Testing register...")

	var isConnexionSuccess = false
	var err error

	t.Run("Register", func(t *testing.T) {
		if err := Register(); err == nil {
			isConnexionSuccess = true
		}
	})

	fmt.Println("Testing connect...")

	t.Run("Connect", func(t *testing.T) {
		if err := Connect(); err == nil {
			isConnexionSuccess = true
		}
	})

	if !isConnexionSuccess {
		t.Fatalf("Unable to connect to the server: %v", err)
		return
	}

	fmt.Println("Connexion success...")
	fmt.Println("Testing Status...")

	t.Run("Status", func(t *testing.T) {
		if err := Status(); err == nil {
			isConnexionSuccess = true
		}
	})

	fmt.Println("Status success...")
	fmt.Println("Testing create entity...")

	t.Run("CreateEntity", func(t *testing.T) {
		if err := CreateEntity(); err != nil {
			t.Fatalf("CreateEntity failed: %v", err)
		}
	})

	fmt.Println("Testing GetEntities...")

	t.Run("GetEntities", func(t *testing.T) {
		if err := GetEntities(); err != nil {
			t.Fatalf("GetEntities failed: %v", err)
		}
	})

	fmt.Println("Testing NewMessage...")

	t.Run("NewMessage", func(t *testing.T) {
		if err := NewMessage(); err != nil {
			t.Fatalf("NewMessage failed: %v", err)
		}
	})

	fmt.Println("Testing MakeDecision...")

	t.Run("MakeDecision", func(t *testing.T) {
		if err := MakeDecision(); err != nil {
			t.Fatalf("MakeDecision failed: %v", err)
		}
	})

	fmt.Println("Testing RemoveEntity...")

	t.Run("RemoveEntity", func(t *testing.T) {
		if err := RemoveEntity(); err != nil {
			t.Fatalf("RemoveEntity failed: %v", err)
		}
	})

	fmt.Println("Testing RemoveUser...")
	t.Run("RemoveUser", func(t *testing.T) {
		if err := RemoveUser(); err != nil {
			t.Fatalf("RemoveUser failed: %v", err)
		}
	})
}
