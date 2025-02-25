package websocketServices

import (
	"fmt"
	"my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/pkg"
	"regexp"
	"strings"

	"github.com/gorilla/websocket"
)

func TalkToWebSocket(token string, message string, interlocutor string) (string, error) {
	UserId, err := pkg.GetUserIDFromJWT(token)

	if err != nil {
		return "error during the process", err
	}

	_, err_ := services.IsExistById(UserId)

	if err_ != nil {
		return "error during the process", err
	}

	prompt, err_ := services.GetPromptByID(UserId)

	if err_ != nil {
		return "error during the process", err
	}

	back, err := services.GptTalkToRequest(message, prompt, interlocutor)
	if err != nil {
		return "error during the process", err
	}

	fmt.Println("Received from GPT: " + back)

	re := regexp.MustCompile(`Response:\s*(.*)`)
	match := re.FindStringSubmatch(back)

	if len(match) > 1 {
		response := match[1]
		return response, nil
	} else {
		fmt.Println("Response non trouvé")
		return "error during the process", fmt.Errorf("error during the process")
	}
}

func MakeDecisionWebSocket(conn *websocket.Conn, msg websocketModels.MakeDecisionRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "MakeDecision"

	back, err := services.GptSimpleRequest(msg.Message)
	if err != nil {
		println("Error in MakeDecisionWebSocket: " + err.Error())
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while calling MakeDecision service",
		})
		return
	}

	if strings.Contains(back, "TalkTo:") {
		re := regexp.MustCompile(`(?m)^TalkTo:\s*(.+)`)
		match := re.FindStringSubmatch(back)

		if len(match) > 1 {
			entity := match[1]

			// ATTENTION: Changer la string envoyée à talkto
			// Il faut implémenter dans la fonction talkto, un système qui récupère les anciens messages en bdd
			// et les store quand il en envoie

			message, err := TalkToWebSocket(msg.Token, "Hello World", entity)
			if err != nil {
				sendError(conn, initialRoute, map[string]interface{}{
					"message": "Error while calling MakeDecision service",
				})
				return
			} else {
				sendResponse(conn, initialRoute, map[string]interface{}{
					"message": "TalkTo: " + entity + "\nMessage: " + message,
				})
				return
			}
		}
	}

	sendResponse(conn, initialRoute, map[string]interface{}{
		"message": back,
	})
}
