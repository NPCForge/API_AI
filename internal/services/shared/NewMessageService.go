package sharedServices

import "my-api/internal/services"

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

		idCurentEntitySending, err := services.GetEntityIdByChecksum(checksumCurrentEntityReceiving)

		if err != nil {
			return err
		}

		_, err = services.NewMessage(idEntitySending, idCurentEntitySending, Message)

		if err != nil {
			return err
		}
	}

	return nil
}
