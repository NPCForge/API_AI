package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"

	"my-api/config"
	httpModels "my-api/internal/models/http"
)

// Function to read the content of a file
func readPromptFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading the file: %w", err)
	}
	return string(content), nil
}

func GptSimpleRequest(userPrompt string, systemPrompt string) (string, error) {
	GptClient := resty.New()

	// Prepare messages with a system prompt
	var Messages = []httpModels.ChatGptSimpleRequestBodyMessage{
		{
			Role:    "system",
			Content: systemPrompt,
		},
		{
			Role:    "user",
			Content: userPrompt,
		},
	}

	//pkg.DisplayContext("GPTTalkToRequest: \n"+
	//	"UserPrompt = \n{\n"+userPrompt+"\n}\n"+
	//	"SystemPrompt = \n{\n"+systemPrompt+"\n}\n",
	//	pkg.Debug,
	//)

	// Create the request body
	body := httpModels.ChatGptSimpleRequestBody{
		Model:    "gpt-3.5-turbo",
		Messages: Messages,
	}

	// Make the request to the OpenAI API
	resp, err := GptClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+config.GetEnvVariable("CHATGPT_TOKEN")).
		SetBody(body).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		return "", fmt.Errorf("error during the request: %w", err)
	}

	// Process the response
	var response httpModels.ChatGPTResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return "", fmt.Errorf("error decoding the response: %w", err)
	}

	// Check if there are valid choices in the response
	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("[GptSimpleRequest]: no response available")
}
