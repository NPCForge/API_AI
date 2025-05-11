package helpers

import (
	"fmt"
	"my-api/internal/services"
	"my-api/pkg"
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

func GetAllDiscussions(EntityChecksum string) (string, error) {
	discussions, err := services.GetDiscussions(EntityChecksum)

	if err != nil {
		pkg.DisplayContext("Cannot retrieve discussions", pkg.Error, err)
		return "", err
	}

	var allDiscussions strings.Builder

	for _, msg := range discussions {
		var SenderChecksum string

		if msg.SenderChecksum == "You" {
			SenderChecksum = EntityChecksum
		} else {
			SenderChecksum = msg.SenderChecksum
		}

		name, err := services.GetEntityNameByChecksum(SenderChecksum)
		if err != nil {
			return "", err
		}
		allDiscussions.WriteString(fmt.Sprintf("[%s -> %s: %s], ", name, msg.ReceiverChecksums, msg.Message))
	}

	allDiscussions.WriteString("\n")

	return allDiscussions.String(), nil
}
