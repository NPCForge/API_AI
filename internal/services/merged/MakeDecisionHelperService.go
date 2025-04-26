package service

import (
	"my-api/internal/services"
	"strconv"
	"strings"
)

func IsMyEntity(Checksum string, Id string) (bool, error) {
	id_entity_owner, err := services.GetEntitiesOwnerByChecksum(Checksum)
	if err != nil {
		return false, err
	}

	if Id == strconv.Itoa(id_entity_owner) {
		return true, nil
	}
	return false, nil
}

func NeedToFinish(msg string) bool {
	for _, str := range strings.Fields(msg) {
		if str == "end_of_discussion" {
			return true
		}
	}
	return false
}
