package services

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"my-api/config"
	"my-api/internal/models/http"
)

func GptSimpleRequest(message string) (string, error) {
	GptClient := resty.New()

	var Messages = []httpModels.ChatGptSimpleRequestBodyMessage{ // Utilise un slice ici
		{
			Role:    "user",
			Content: message,
		},
	}

	var body httpModels.ChatGptSimpleRequestBody = httpModels.ChatGptSimpleRequestBody{
		Model:    "gpt-3.5-turbo",
		Messages: Messages, // Envoie le slice de messages
	}

	resp, err := GptClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+config.GetEnvVariable("CHATGPT_TOKEN")).
		SetBody(body).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		return "", fmt.Errorf("erreur lors de la requête : %w", err)
	}

	var response httpModels.ChatGPTResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return "", fmt.Errorf("erreur lors du déchiffrage de la réponse : %w", err)
	}
	// Vérifie s'il y a des choix dans la réponse
	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("aucune réponse disponible")
}
