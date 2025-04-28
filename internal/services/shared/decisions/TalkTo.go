package decisions

import (
	"fmt"
	"my-api/internal/services"
	"my-api/internal/services/helpers"
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

	roleDescription, err := services.GetPromptByID(strconv.Itoa(EntityId))

	if err != nil {
		pkg.DisplayContext("Cannot get prompt using entity ID", pkg.Error, err)
		return "", err, false
	}

	systemPrompt, err := services.ReadPromptFromFile("prompts/discussion.txt")
	if err != nil {
		return "", fmt.Errorf("error retrieving the system prompt: %w", err), false
	}

	systemPrompt = strings.Replace(systemPrompt, "{Role Description Here}", roleDescription, 1)

	userPrompt := "Interlocutor: " + interlocutorChecksum + "\nDiscussion: { " + message + " }"

	back, err := services.GptSimpleRequest(userPrompt, systemPrompt)

	if helpers.NeedToFinish(back) {
		pkg.DisplayContext("After this, need to finish", pkg.Debug)
	}

	re := regexp.MustCompile(`Response:\s*(.*)`)
	match := re.FindStringSubmatch(back)

	if len(match) > 1 {
		response := match[1]
		//pkg.DisplayContext("Expected response: "+response, pkg.Debug)
		return response, nil, helpers.NeedToFinish(back)
	} else {
		pkg.DisplayContext("userPrompt = "+userPrompt, pkg.Debug)
		pkg.DisplayContext("systemPrompt = "+systemPrompt, pkg.Debug)
		pkg.DisplayContext("Response pattern not found in GPT output: "+back, pkg.Error, true)
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

func HandleTalkToLogic(Decision string, Checksum string) (string, error) {
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
