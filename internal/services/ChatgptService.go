package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go-resty/resty/v2"

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

func GptTalkToRequest(message string, prompt string, interlocutor string) (string, error) {
	GptClient := resty.New()

	systemPrompt, err := readPromptFromFile("prompts/discussion.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err)
	}

	systemPrompt += "\n" + prompt

	userPrompt := "Interlocutor: " + interlocutor + "\nDiscussion: { " + message + " }"

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

	yellow := "\033[33m"
	reset := "\033[0m"

	println("GPTTalkToRequest: \n",
		string(yellow), "userPrompt = {\n", userPrompt, "\n}\n",
		string(reset),
	)

	// Create the request body
	body := httpModels.ChatGptSimpleRequestBody{
		Model:    "gpt-3.5-turbo",
		Messages: Messages,
	}

	// body := httpModels.LmStudioSimpleRequestBody{
	// 	Model:       "llama-3.2-3b-instruct",
	// 	Messages:    Messages,
	// 	Temperature: 0.5,
	// 	Max_tokens:  -1,
	// 	Stream:      false,
	// }

	// Make the request to the OpenAI API
	resp, err := GptClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+config.GetEnvVariable("CHATGPT_TOKEN")).
		SetBody(body).
		Post("https://api.openai.com/v1/chat/completions")

	// resp, err := GptClient.R().
	// 	SetHeader("Content-Type", "application/json").
	// 	SetBody(body).
	// 	Post("http://localhost:1234/v1/chat/completions")

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

	return "", fmt.Errorf("[GptTalkToRequest]: no response available")
}

func GptSimpleRequest(message string) (string, error) {
	GptClient := resty.New()

	// Read the "curriculum.txt" file to get the system prompt
	systemPrompt, err := readPromptFromFile("prompts/curriculum.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err)
	}

	// Prepare messages with a system prompt
	var Messages = []httpModels.ChatGptSimpleRequestBodyMessage{
		{
			Role:    "system",
			Content: systemPrompt,
		},
		{
			Role:    "user",
			Content: message,
		},
	}

	yellow := "\033[33m"
	reset := "\033[0m"

	println("GPTTalkToRequest: \n",
		string(yellow), "userPrompt = {\n", message, "\n}\n",
		string(reset),
	)

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
