package apiTesting

import (
	"encoding/json"
	"os"
)

func SaveToken(kind string, token string) error {
	tokens := make(map[string]string)

	file, err := os.ReadFile(TokenFile)
	if err == nil {
		_ = json.Unmarshal(file, &tokens)
	}

	tokens[kind] = token

	jsonData, _ := json.MarshalIndent(tokens, "", "  ")
	return os.WriteFile(TokenFile, jsonData, 0644)
}

func ResetTokenFile() error {
	return os.WriteFile(TokenFile, []byte("{}"), 0644)
}
