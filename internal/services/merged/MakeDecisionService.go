package service

import (
	"errors"
	"fmt"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"regexp"
	"strconv"
	"strings"
)

func talkTo(Checksum string, message string, interlocutorChecksum string) (string, error, bool) {
	EntityId, err := services.GetEntityIdByChecksum(Checksum)

	if err != nil {
		pkg.DisplayContext("Cannot get Entity ID using checksum", pkg.Error, err)
		return "", err, false
	}

	prompt, err := services.GetPromptByID(strconv.Itoa(EntityId))

	if err != nil {
		pkg.DisplayContext("Cannot get prompt using entity ID", pkg.Error, err)
		return "", err, false
	}

	systemPrompt, err := readPromptFromFile("prompts/discussion.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err), false
	}

	systemPrompt += "\n" + prompt
	userPrompt := "Interlocutor: " + interlocutorChecksum + "\nDiscussion: { " + message + " }"

	back, err := GptSimpleRequest(userPrompt, systemPrompt)

	pkg.DisplayContext("TalkTo GPT request = "+back, pkg.Debug)

	if NeedToFinish(back) {
		pkg.DisplayContext("After this, need to finish", pkg.Debug)
	}

	re := regexp.MustCompile(`Response:\s*(.*)`)
	match := re.FindStringSubmatch(back)

	if len(match) > 1 {
		response := match[1]
		pkg.DisplayContext("Expected response: "+response, pkg.Debug)
		return response, nil, NeedToFinish(back)
	} else {
		pkg.DisplayContext("Response pattern not found in GPT output: "+back, pkg.Error)
		return "", fmt.Errorf("error during the process"), false
	}
}

func getAllDiscussionsForEntity(EntityChecksum string, InterlocutorChecksums []string) (string, error) {
	EntityID, err := services.GetEntityIdByChecksum(EntityChecksum)

	if err != nil {
		pkg.DisplayContext("Cannot get Entity ID using checksum", pkg.Error, err)
		return "", err
	}

	var allDiscussions string

	for _, checksum := range InterlocutorChecksums {
		interlocutorID, err := services.GetEntityIdByChecksum(checksum)

		if err != nil {
			pkg.DisplayContext("Cannot get Interlocutor ID using checksum: "+checksum, pkg.Error, err)
			return "", err
		}

		discussion, err := services.GetDiscussion(strconv.Itoa(EntityID), strconv.Itoa(interlocutorID))

		if err != nil {
			pkg.DisplayContext("Cannot get Discussion", pkg.Error, err)
			return "", err
		}

		var sb strings.Builder

		for _, msg := range discussion {
			sb.WriteString(fmt.Sprintf("[%s -> %s: %s], ",
				msg.SenderChecksum, msg.ReceiverChecksums, msg.Message))
		}

		allDiscussions += sb.String()
		allDiscussions += "\n"
	}

	return allDiscussions, nil
}

func askLLMForDecision(Message string, Checksum string) (string, error) {
	// Get raw new messages from db to format them
	newMessages, err := services.GetNewMessages(Checksum)

	if err != nil {
		return "", err
	}

	Message += "\nNew Messages: {" + strings.Join(newMessages, ", ") + "}"

	// Read the "curriculum.txt" file to get the system prompt
	systemPrompt, err := readPromptFromFile("prompts/curriculum.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err)
	}

	decision, err := GptSimpleRequest(Message, systemPrompt)
	if err != nil {
		pkg.DisplayContext("GptSimpleRequest failed: ", pkg.Error, err)
		return "", err
	}

	return decision, nil
}

func handleTalkToLogic(Decision string, Checksum string) (string, error) {
	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindStringSubmatch(Decision)

	if len(matches) > 1 {
		checksums := strings.Split(matches[1], ", ")
		checksumsString := "[" + strings.Join(checksums, ", ") + "]"

		discussions, err := getAllDiscussionsForEntity(Checksum, checksums)

		if err != nil {
			pkg.DisplayContext("Cannot get Discussions using checksum", pkg.Error, err)
			return "", err
		}

		message, err, _ := talkTo(Checksum, discussions, checksumsString) // should finish is not used. Must implement

		if err != nil {
			pkg.DisplayContext("TalkToPreprocess failed: ", pkg.Error, err)
			return "", err
		}
		return "TalkTo: " + checksumsString + "\nMessage: " + message, nil
	}
	return "", fmt.Errorf("Cannot find interlocutor checksums: " + Decision)
}

func interpretLLMDecision(Decision string, Checksum string) (string, error) {
	if strings.Contains(Decision, "TalkTo:") {
		return handleTalkToLogic(Decision, Checksum)
	}

	return "", fmt.Errorf("no such Decision: %s", Decision)
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
