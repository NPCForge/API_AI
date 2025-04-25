package service

import (
	"errors"
	"my-api/internal/services"
	"my-api/internal/utils"
	"strconv"
)

func isMyEntity(Checksum string, Id string) (bool, error) {
	id_entity_owner, err := services.GetEntitiesOwnerByChecksum(Checksum)
	if err != nil {
		return false, err
	}

	if Id == strconv.Itoa(id_entity_owner) {
		return true, nil
	}
	return false, nil
}

func MakeDecisionService(Message string, Checksum string, Token string) (string, error) {
	id, err := utils.GetUserIDFromJWT(Token)

	if err != nil {
		return "", err
	}

	// verifie si l'entité appartient a l'utilisateur
	val, err := isMyEntity(Checksum, id)

	if err != nil {
		return "", err
	}

	if !val {
		return "", errors.New("access denied to this entity")
	}

	// recupere les messages non lue adresser à l'entity
	// newMessages, _ := services.GetNewMessages(receiver)

	// donner a gpt

	// il donne une liste de message formater par destinataire

	// je renvoie ça formater

	return id + " ", nil
}
