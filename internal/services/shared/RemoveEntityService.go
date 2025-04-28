package sharedServices

import (
	"fmt"
	"my-api/internal/services"
	"my-api/internal/utils"
	"strconv"
)

// RemoveEntityService deletes an entity if the requesting user is the verified owner.
func RemoveEntityService(checksum string, self string) error {
	id, err := services.GetEntitiesOwnerByChecksum(checksum)
	if err != nil {
		return err
	}

	selfID, err := utils.GetUserIDFromJWT(self)
	if err != nil {
		return err
	}

	selfInt, err := strconv.Atoi(selfID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	if id != selfInt {
		return fmt.Errorf("access denied: you are not the owner of this entity")
	}

	return services.DropEntityByChecksum(checksum)
}
