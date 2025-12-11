package apiTesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Reusing struct definition locally for test
type EntityCheck struct {
	Checksum string `json:"checksum"`
}

type GetEntitiesResponseCheck struct {
	Entities []EntityCheck `json:"entities"`
}

func EntityDied() error {
	token, err := ReadTokens("http")
	if err != nil {
		return err
	}

	// 1. Create a victim entity
	fmt.Println("Creating a victim entity...")
	createReq := map[string]string{
		"action":   "Create",
		"name":     "Victim",
		"prompt":   "I am a victim",
		"checksum": "victim_checksum",
		"role":     "VICTIM",
		"token":    token,
	}
	payload, _ := json.Marshal(createReq)
	resp, err := http.Post(HttpBaseUrl+"CreateEntity", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create victim: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to create victim, status: %d", resp.StatusCode)
	}

	// 2. Kill the victim
	fmt.Println("Killing the victim...")
	diedReq := map[string]string{
		"checksum": "victim_checksum",
	}
	payloadDied, _ := json.Marshal(diedReq)

	req, err := http.NewRequest("POST", HttpBaseUrl+"EntityDied", bytes.NewBuffer(payloadDied))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	respDied, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call EntityDied: %w", err)
	}
	defer respDied.Body.Close()

	if respDied.StatusCode != 200 {
		return fmt.Errorf("EntityDied failed with status: %d", respDied.StatusCode)
	}

	// 3. Verify it's gone using GetEntities
	fmt.Println("Verifying victim is gone...")
	reqGet, err := http.NewRequest("GET", HttpBaseUrl+"GetEntities", nil)
	if err != nil {
		return err
	}
	reqGet.Header.Set("Authorization", token)
	reqGet.Header.Set("Content-Type", "application/json")

	respGet, err := client.Do(reqGet)
	if err != nil {
		return fmt.Errorf("failed to call GetEntities: %w", err)
	}
	defer respGet.Body.Close()

	if respGet.StatusCode != 200 {
		return fmt.Errorf("GetEntities failed with status: %d", respGet.StatusCode)
	}

	var data GetEntitiesResponseCheck
	if err := json.NewDecoder(respGet.Body).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode GetEntities response: %w", err)
	}

	for _, e := range data.Entities {
		if e.Checksum == "victim_checksum" {
			return fmt.Errorf("victim entity still exists after EntityDied")
		}
	}

	return nil
}
