package service

import "my-api/internal/utils"

func MakeDecisionService(Message string, Checksum string, Token string) (string, error) {
	receiver, err := utils.GetUserIDFromJWT(Token)

	if err != nil {
		return "", err
	}

	// recupere les messages non lue adresser à l'entity
	// newMessages, _ := services.GetNewMessages(receiver)

	return receiver + " ", nil
}
