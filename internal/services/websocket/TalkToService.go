package websocketServices

import (
	"fmt"
	"my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/pkg"
	"regexp"

	"github.com/gorilla/websocket"
)

func TalkToWebSocket(conn *websocket.Conn, msg websocketModels.TalkToRequest, sendResponse func(*websocket.Conn, string, map[string]interface{}), sendError func(*websocket.Conn, string, map[string]interface{})) {
	var initialRoute = "TalkTo"

	UserId, err := pkg.GetUserIDFromJWT(msg.Token)

	if err != nil {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "error during the process",
		})
	}

	_, err_ := services.IsExistById(UserId)

	if err_ != nil {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "failed",
		})
		return
	}

	prompt, err_ := services.GetPromptByID(UserId)

	if err_ != nil {
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": "failed",
		})
		return
	}

	back, err := services.GptTalkToRequest(msg.Message, prompt, msg.Interlocutor)
	if err != nil {
		println("Error in TalkToWebSocket: " + err.Error())
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Error while calling TalkTo service",
		})
		return
	}

	fmt.Println("Received from GPT: " + back)

	re := regexp.MustCompile(`Response:\s*(.*)`)
	match := re.FindStringSubmatch(back)

	if len(match) > 1 {
		response := match[1]
		sendResponse(conn, initialRoute, map[string]interface{}{
			"message": response,
		})
	} else {
		fmt.Println("Response non trouvé")
	}

}
