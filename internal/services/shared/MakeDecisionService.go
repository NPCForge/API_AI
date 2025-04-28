package sharedServices

import (
	"errors"
	"fmt"
	"my-api/internal/services"
	"my-api/internal/services/shared/decisions"
	"my-api/internal/utils"
	"my-api/pkg"
	"strings"
)

func interpretLLMDecision(Decision string, Checksum string) (string, error) {
	if strings.Contains(Decision, "TalkTo:") {
		return decisions.HandleTalkToLogic(Decision, Checksum)
	}

	return "", fmt.Errorf("no such Decision: %s", Decision)
}

func askLLMForDecision(Message string, Checksum string) (string, error) {
	newMessages, err := services.GetNewMessages(Checksum)

	if err != nil {
		return "", err
	}

	Message += "\nNew Messages: {" + strings.Join(newMessages, ", ") + "}"

	// Read the "curriculum.txt" file to get the system prompt
	systemPrompt, err := services.ReadPromptFromFile("prompts/curriculum.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err)
	}

	//pkg.DisplayContext("userPrompt = "+Message, pkg.Debug)
	//pkg.DisplayContext("systemPrompt = "+systemPrompt, pkg.Debug)

	decision, err := services.GptSimpleRequest(Message, systemPrompt)
	if err != nil {
		pkg.DisplayContext("GptSimpleRequest failed: ", pkg.Error, err)
		return "", err
	}

	return decision, nil
}

func MakeDecisionService(Message string, Checksum string, Token string) (string, error) {
	id, err := utils.GetUserIDFromJWT(Token)

	if err != nil {
		return "", err
	}

	val, err := IsMyEntity(Checksum, id)

	if err != nil {
		return "", err
	}

	if !val {
		return "", errors.New("access denied to this entity")
	}

	decision, err := askLLMForDecision(Message, Checksum)

	if err != nil {
		pkg.DisplayContext("Error after decision making: ", pkg.Error, err)
		return "", err
	}

	task, err := interpretLLMDecision(decision, Checksum)

	if err != nil {
		pkg.DisplayContext("Error after llm response interpretation: ", pkg.Error, err)
		return "", err
	}

	return task, nil
}
