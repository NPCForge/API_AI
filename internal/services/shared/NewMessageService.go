package sharedServices

import "my-api/internal/services"

// NewMessageService creates a new message from one entity to one or multiple receiver entities.
func NewMessageService(
	ChecksumEntitySending string,
	ChecksumEntityReceiving []string,
	Message string,
) error {
	idEntitySending, err := services.GetEntityIdByChecksum(ChecksumEntitySending)
	if err != nil {
		return err
	}

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
