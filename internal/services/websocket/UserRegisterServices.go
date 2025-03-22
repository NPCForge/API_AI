package websocketServices

import (
	"errors"
	"strconv"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"

	websocketModels "my-api/internal/models/websocket"
	"my-api/internal/services"
	"my-api/internal/types"
	"my-api/pkg"
)

func SaveInDatabase(entity websocketModels.RegisterRequest) (int64, error) {
	color.Cyan("📝 SaveInDatabase triggered with checksum: %s", entity.Checksum)

	exists, err := services.IsExist(entity.Checksum)
	if err != nil {
		color.Red("❌ Error checking existence: %v", err)
		return -1, errors.New("error searching in table")
	}

	if exists {
		color.Yellow("⚠️ Entry already exists in DB for checksum: %s", entity.Checksum)
		return -1, errors.New("entry already exists in database")
	}

	id, err := services.RegisterWebsocket(entity.Checksum, entity)
	if err != nil {
		color.Red("❌ Error while registering: %v", err)
		return -1, errors.New("error while registering")
	}

	color.Green("✅ New user registered with ID: %d", id)
	return id, nil
}

func RegisterServiceWebSocket(
	conn *websocket.Conn,
	msg websocketModels.RegisterRequest,
	sendResponse types.SendResponseFunc,
	sendError types.SendErrorFunc,
) {
	initialRoute := "Register"

	id, err := SaveInDatabase(msg)
	if err != nil {
		color.Red("❌ Failed to save in database: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	stringId := strconv.FormatInt(id, 10)

	pass, err := pkg.GenerateJWT(stringId)
	if err != nil {
		color.Red("❌ Failed to generate JWT: %v", err)
		sendError(conn, initialRoute, map[string]interface{}{
			"message": "Unable to generate Token",
		})
		return
	}

	pkg.SetToken(stringId, pass)

	color.Green("✅ Token generated and stored for user: %s", stringId)
	sendResponse(conn, initialRoute, map[string]interface{}{
		"token": pass,
	})
}
