package apiTesting

import (
	"encoding/json"
	"errors"
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

func ReadTokens(connType string) (string, error) {
	file, err := os.ReadFile(TokenFile)
	if err != nil {
		return "", err
	}

	var tokens map[string]string
	if err := json.Unmarshal(file, &tokens); err != nil {
		return "", err
	}

	token := tokens[connType]

	if token == "" {
		return "", errors.New("token not found")
	}

	return token, nil
}
