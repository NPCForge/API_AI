package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	sharedModels "my-api/internal/models/shared"
	"my-api/pkg"

	"github.com/go-resty/resty/v2"
)

// ReadPromptFromFile reads the content of a file and returns it as a string.
func ReadPromptFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading the file: %w", err)
	}
	return string(content), nil
}

// GptSimpleRequest sends a user prompt and system prompt to the OpenAI API and returns the model's response.
func GptSimpleRequest(userPrompt string, systemPrompt string) (string, error) {
	GptClient := resty.New()

	//pkg.DisplayContext("SystemPrompt = "+systemPrompt, pkg.Debug)
	//pkg.DisplayContext("userPrompt = "+userPrompt, pkg.Debug)

	// Prepare the chat messages
	messages := []sharedModels.ChatGptSimpleRequestBodyMessage{
		{
			Role:    "system",
			Content: systemPrompt,
		},
		{
			Role:    "user",
			Content: userPrompt,
		},
	}

	// Create the request body
	body := sharedModels.ChatGptSimpleRequestBody{
		Model:    "llama3.1:70b",
		Messages: messages,
		Stream:   false,
	}

	// Send the request to the OpenAI API
	resp, err := GptClient.R().
		SetHeader("Content-Type", "application/json").
		//SetHeader("Authorization", "Bearer "+config.GetEnvVariable("CHATGPT_TOKEN")).
		SetBody(body).
		Post("https://bzcbmeimpcqpvd-11434.proxy.runpod.net/v1/chat/completions")

	if err != nil {
		return "", fmt.Errorf("error during the request: %w", err)
	}

	// Process the response
	var response sharedModels.ChatGPTResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return "", fmt.Errorf("error decoding the response: %w", err)
	}

	// Check if there are valid choices in the response
	if len(response.Choices) > 0 {
		pkg.DisplayContext("Response = "+response.Choices[0].Message.Content, pkg.Debug)
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("[GptSimpleRequest]: no response available")
}
