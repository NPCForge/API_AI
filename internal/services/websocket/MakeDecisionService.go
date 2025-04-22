package websocketServices

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"

	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/internal/types"
	"my-api/internal/utils"
)

func NeedToFinish(msg string) bool {
	for _, str := range strings.Fields(msg) {
		if str == "end_of_discussion" {
			return true
		}
	}
	return false
}

func TalkToWebSocket(token string, message string, interlocutor string) (string, error, bool) {
	UserId, err := utils.GetUserIDFromJWT(token)
	if err != nil {
		color.Red("❌ JWT parsing failed: %v", err)
		return "error during the process", err, false
	}

	_, err_ := services.IsExistById(UserId)
	if err_ != nil {
		color.Red("❌ User ID doesn't exist: %v", err_)
		return "error during the process", err_, false
	}

	prompt, err_ := services.GetPromptByID(UserId)
	if err_ != nil {
		color.Red("❌ Failed to get prompt: %v", err_)
		return "error during the process", err_, false
	}

	back, err := services.GptTalkToRequest(message, prompt, interlocutor)
	if err != nil {
		color.Red("❌ GPT TalkToRequest failed: %v", err)
		return "error during the process", err, false
	}

	color.Cyan("📥 Received from GPT (%d): %s", len(back), back)

	if NeedToFinish(back) {
		color.HiMagenta("𝌦 After this, we need to finish : %s", back)
	}

	re := regexp.MustCompile(`Response:\s*(.*)`)
	match := re.FindStringSubmatch(back)

	if len(match) > 1 {
		response := match[1]
		color.Green("✅ Extracted Response: %s", response)
		return response, nil, NeedToFinish(back)
	} else {
		color.Yellow("⚠️ Response pattern not found in GPT output")
		return "error during the process", fmt.Errorf("error during the process"), false
	}
}

func TalkToPreprocess(msg websocketModels.MakeDecisionRequest, entity_names []string) (string, error, bool) {
	from, err := utils.GetUserIDFromJWT(msg.Token)

	if err != nil {
		color.Red("❌ JWT parsing failed in TalkToPreprocess: %v", err)
		return "Error getting user ID", err, false
	}

	var allDiscussions string

	for _, name := range entity_names {
		to, err := services.GetEntityByName(name)

		if err != nil {
			color.Red("❌ Entity '%s' not found: %v", to, err)
			return "error during the process", err, false
		}

		discussion, err := services.GetDiscussion(from, to)

		if err != nil {
			color.Red("❌ Failed to retrieve discussion: %v", err)
			return "error during the process", err, false
		}

		var sb strings.Builder

		for _, msg := range discussion {
			sb.WriteString(fmt.Sprintf("[%s -> %s: %s], ",
				msg.SenderName, msg.ReceiverNames, msg.Message))
		}

		allDiscussions += sb.String()
		allDiscussions += "\n"
	}

	color.Cyan("💬 Compiled discussion:\n%s", allDiscussions)

	namesString := "[" + strings.Join(entity_names, ", ") + "]"

	message, err, finish := TalkToWebSocket(msg.Token, allDiscussions, namesString)

	if err != nil {
		color.Red("❌ TalkToWebSocket failed: %v", err)
		return "error during the process", err, false
	}

	return message, nil, finish
}

func MakeDecisionWebSocket(
	conn *websocket.Conn,
	msg websocketModels.MakeDecisionRequest,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) {
	var initialRoute = "MakeDecision"
	receiver, err := utils.GetUserIDFromJWT(msg.Token)

	if err != nil {
		color.Red("❌ JWT parsing failed: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Invalid token",
		})
		return
	}

	color.Green("🧾 Receiver from JWT: %v", receiver)

	newMessages, _ := services.GetNewMessages(receiver)

	var formattedMessages []string
	for _, msg := range newMessages {
		formattedMessages = append(formattedMessages, fmt.Sprintf("[%s -> %s: \"%s\"]", msg.SenderName, msg.ReceiverNames, msg.Message))
	}

	result := "New Messages: {" + strings.Join(formattedMessages, ", ") + "}"
	color.Cyan("🚀 New Messages: %s", result)
	color.Yellow("🌿 Raw messages: %+v", newMessages)

	msg.Message += "\n" + result

	back, err := services.GptSimpleRequest(msg.Message)
	if err != nil {
		color.Red("❌ GptSimpleRequest failed: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while calling MakeDecision service",
		})
		return
	}

	if strings.Contains(back, "TalkTo:") {

		re := regexp.MustCompile(`\[(.*?)\]`)
		matches := re.FindStringSubmatch(back)

		if len(matches) > 1 {
			names := strings.Split(matches[1], ", ")
			namesString := "[" + strings.Join(names, ", ") + "]"

			message, err, _ := TalkToPreprocess(msg, names)

			if err != nil {
				color.Red("❌ Error during TalkToPreprocess: %v", err)
				sendError(conn, initialRoute, map[string]interface{}{
					"message": "Error while calling MakeDecision service",
				})
				return
			} else {
				sendResponse(conn, initialRoute, map[string]interface{}{
					"message": "TalkTo: " + namesString + "\nMessage: " + message,
				})
				return
			}
		}

		//match := re.FindStringSubmatch(back)
		//
		//if len(match) > 1 {
		//	entity := match[1]
		//
		//	message, err := TalkToPreprocess(msg, entity)
		//	if err != nil {
		//		sendError(conn, initialRoute, map[string]interface{}{
		//			"message": "Error while calling MakeDecision service",
		//		})
		//		return
		//	} else {
		//		sendResponse(conn, initialRoute, map[string]interface{}{
		//			"message": "TalkTo: " + entity + "\nMessage: " + message,
		//		})
		//		return
		//	}
		//}
	} else {
		println("[MakeDecisionWebSocket]: Unable to find TalkTo")
	}
}
