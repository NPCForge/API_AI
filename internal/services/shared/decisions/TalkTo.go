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

// talkTo generates a message for an entity to communicate with a specific interlocutor using GPT based on their checksums.
func talkTo(Checksum string, message string, interlocutorChecksum string, GamePrompt string, Role string) (string, error, bool) {
	EntityId, err := services.GetEntityIdByChecksum(Checksum)
	if err != nil {
		pkg.DisplayContext("Cannot get Entity ID using checksum", pkg.Error, err)
		return "", err, false
	}

	characterDescription, err := services.GetPromptByID(strconv.Itoa(EntityId))
	if err != nil {
		pkg.DisplayContext("Cannot get prompt using entity ID", pkg.Error, err)
		return "", err, false
	}

	systemPrompt, err := services.ReadPromptFromFile("prompts/Talk.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err), false
	}

	systemPrompt = strings.Replace(systemPrompt, "{Personality Description Here}", characterDescription, 1)
	systemPrompt = strings.Replace(systemPrompt, "{Role Description Here}", Role, 1)
	systemPrompt = strings.Replace(systemPrompt, "{Game Prompt Here}", GamePrompt, 1)

	userPrompt := "Discussion: { " + message + " }"

	back, err := services.GptSimpleRequest(userPrompt, systemPrompt)
	if helpers.NeedToFinish(back) {
		pkg.DisplayContext("Conversation marked to be finished", pkg.Debug)
	}

	var data map[string]string

	err = json.Unmarshal([]byte(back), &data)
	if err != nil {
		pkg.DisplayContext("Cannot unmarshal gpt response data:", pkg.Error, err)
		return "", err, false
	}

	if data["Response"] == "" {
		return talkTo(Checksum, message, interlocutorChecksum, GamePrompt, Role)
	}

	return data["Response"], nil, false
}

// getAllDiscussionsForEntity retrieves all discussions for an entity against a list of interlocutors.
func getAllDiscussionsForEntity(EntityChecksum string, InterlocutorChecksums []string) (string, error) {
	EntityID, err := services.GetEntityIdByChecksum(EntityChecksum)
	if err != nil {
		pkg.DisplayContext("Cannot get Entity ID using checksum", pkg.Error, err)
		return "", err
	}

	var allDiscussions strings.Builder

	for _, checksum := range InterlocutorChecksums {
		interlocutorID, err := services.GetEntityIdByChecksum(checksum)
		if err != nil {
			pkg.DisplayContext("Cannot get Interlocutor ID using checksum: "+checksum, pkg.Error, err)
			return "", err
		}

		discussion, err := services.GetDiscussionFromBy(strconv.Itoa(EntityID), strconv.Itoa(interlocutorID))
		if err != nil {
			pkg.DisplayContext("Cannot retrieve discussion", pkg.Error, err)
			return "", err
		}

		for _, msg := range discussion {
			allDiscussions.WriteString(fmt.Sprintf("[%s -> %s: %s], ", msg.SenderChecksum, msg.ReceiverChecksums, msg.Message))
		}
		allDiscussions.WriteString("\n")
	}

	return allDiscussions.String(), nil
}

// HandleTalkToLogic parses the decision string, gathers discussions, and generates a response for the entity to speak to the interlocutors.
func HandleTalkToLogic(Checksum string, GamePrompt string, Role string) (string, error) {
	discussions, err := helpers.GetAllDiscussions(Checksum)
	if err != nil {
		pkg.DisplayContext("Cannot get discussions using checksum", pkg.Error, err)
		return "", err
	}

	message, err, _ := talkTo(Checksum, discussions, "[Everyone]", GamePrompt, Role) // shouldFinish flag not used yet
	if err != nil {
		pkg.DisplayContext("TalkToPreprocess failed:", pkg.Error, err)
		return "", err
	}
	return `{"Action": "TalkTo", "TalkTo": "[Everyone]", "Message": "` + message + `"}`, nil
}
