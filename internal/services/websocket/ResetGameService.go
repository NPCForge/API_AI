package websocketServices

import (
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/internal/types"
)

func ResetGameServiceWebSocket(conn *websocket.Conn,
	msg websocketModels.ResetGameRequest,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc) {
	initialRoute := "ResetGame"
	color.Cyan("🗑️  ResetGameServiceWebSocket triggered")

	err := services.ResetGame()

	if err != nil {
		color.Red(err.Error())
	} else {
		color.Green("🗑️  ResetGameServiceWebSocket success")
		sendResponse(conn, initialRoute, map[string]interface{}{})
	}

}
