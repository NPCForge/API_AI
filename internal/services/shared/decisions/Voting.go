package decisions

import (
	"encoding/json"
	"fmt"
	"my-api/internal/services"
	"my-api/internal/services/helpers"
	"my-api/pkg"
	"strconv"
	"strings"
)

func voteFor(Checksum string, Discussion string, GamePrompt string) (string, error) {
	EntityId, err := services.GetEntityIdByChecksum(Checksum)
	if err != nil {
		pkg.DisplayContext("Cannot get Entity ID using checksum", pkg.Error, err)
		return "", err
	}

	roleDescription, err := services.GetPromptByID(strconv.Itoa(EntityId))
	if err != nil {
		pkg.DisplayContext("Cannot get prompt using entity ID", pkg.Error, err)
		return "", err
	}

	systemPrompt, err := services.ReadPromptFromFile("prompts/Vote.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err)
	}

	systemPrompt = strings.Replace(systemPrompt, "{Role Description Here}", roleDescription, 1)
	systemPrompt = strings.Replace(systemPrompt, "{Game Prompt Here}", GamePrompt, 1)

	userPrompt := "Discussion: { " + Discussion + " }"

	back, err := services.GptSimpleRequest(userPrompt, systemPrompt)

	var data map[string]string

	err = json.Unmarshal([]byte(back), &data)
	if err != nil {
		pkg.DisplayContext("Cannot unmarshal gpt response data:", pkg.Error, err)
		return "", err
	}

	if data["VoteFor"] == "" {
		return voteFor(Checksum, Discussion, GamePrompt)
	} else {
		return data["VoteFor"], nil
	}
}

func HandleVotingLogic(Checksum string, GamePrompt string) (string, error) {
	discussions, err := helpers.GetAllDiscussions(Checksum)
	if err != nil {
		pkg.DisplayContext("Cannot get discussions using checksum", pkg.Error, err)
		return "", err
	}

	name, err := voteFor(Checksum, discussions, GamePrompt)

	if err != nil {
		pkg.DisplayContext("voteFor failed:", pkg.Error, err)
		return "", err
	}
	return `{"Action": "VoteFor", "VoteFor": "` + name + `"}`, nil
}
