package sharedServices

import (
	"my-api/internal/services"
	"my-api/internal/services/gobalHelpers"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"
)

func NewBroadcastMessage(Token string, ChecksumEntitySending string, Message string) error {
	UserID, err := utils.GetUserIDFromJWT(Token)
	if err != nil {
		return err
	}
	sendingID, err := services.GetEntityIdByChecksum(ChecksumEntitySending)
	if err != nil {
		return err
	}

	intUserID, err := strconv.Atoi(UserID)

	_, err = services.BroadcastMessage(intUserID, sendingID, Message)
	if err != nil {
		return err
	}
	return nil
}

// NewMessageService creates a new message from one entity to one or multiple receiver entities.
func NewMessageService(
	ChecksumEntitySending string,
	ChecksumEntityReceiving []string,
	Message string,
	Token string,
) error {
	idEntitySending, err := services.GetEntityIdByChecksum(ChecksumEntitySending)
	if err != nil {
		return err
	}

	if gobalHelpers.StringContains(ChecksumEntityReceiving, "Everyone") != -1 {
		pkg.DisplayContext("Adding broadcast message", pkg.Update)
		return NewBroadcastMessage(Token, ChecksumEntitySending, Message)
	}

	pkg.DisplayContext("Adding private message", pkg.Update)

	for _, checksumCurrentEntityReceiving := range ChecksumEntityReceiving {
		idCurrentEntityReceiving, err := services.GetEntityIdByChecksum(checksumCurrentEntityReceiving)
		if err != nil {
			return err
		}

		_, err = services.NewMessage(idEntitySending, idCurrentEntityReceiving, Message)
		if err != nil {
			return err
		}
	}

	return nil
}
