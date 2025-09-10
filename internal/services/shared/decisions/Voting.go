package decisions

import (
	"context"
	"encoding/json"
	"fmt"
	"my-api/internal/services"
	"my-api/internal/services/helpers"
	"my-api/pkg"
	"strconv"
	"strings"
)

func voteFor(ctx context.Context, Checksum string, Discussion string, GamePrompt string, Role string) (string, error) {
	EntityId, err := services.GetEntityIdByChecksum(Checksum)
	if err != nil {
		pkg.DisplayContext("Cannot get Entity ID using checksum", pkg.Error, err)
		return "", err
	}

	personalityDescription, err := services.GetPromptByID(strconv.Itoa(EntityId))
	if err != nil {
		pkg.DisplayContext("Cannot get prompt using entity ID", pkg.Error, err)
		return "", err
	}

	systemPrompt, err := services.ReadPromptFromFile("prompts/Vote.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err)
	}

	systemPrompt = strings.Replace(systemPrompt, "{Personality Description Here}", personalityDescription, 1)
	systemPrompt = strings.Replace(systemPrompt, "{Role Description Here}", Role, 1)
	systemPrompt = strings.Replace(systemPrompt, "{Game Prompt Here}", GamePrompt, 1)

	userPrompt := "Discussion: { " + Discussion + " }"

	back, err := services.GptSimpleRequest(ctx, userPrompt, systemPrompt)

	var data map[string]string

	err = json.Unmarshal([]byte(back), &data)
	if err != nil {
		pkg.DisplayContext("Cannot unmarshal gpt response data:", pkg.Error, err)
		return "", err
	}

	if data["VoteFor"] == "" {
		return voteFor(ctx, Checksum, Discussion, GamePrompt, Role)
	} else {
		return data["VoteFor"], nil
	}
}

func HandleVotingLogic(ctx context.Context, Checksum string, GamePrompt string, Role string) (string, error) {
	discussions, err := helpers.GetAllDiscussions(Checksum)
	if err != nil {
		pkg.DisplayContext("Cannot get discussions using checksum", pkg.Error, err)
		return "", err
	}

	name, err := voteFor(ctx, Checksum, discussions, GamePrompt, Role)

	if err != nil {
		pkg.DisplayContext("voteFor failed:", pkg.Error, err)
		return "", err
	}
	return `{"Action": "VoteFor", "VoteFor": "` + name + `"}`, nil
}
