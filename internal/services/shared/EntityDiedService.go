package sharedServices

import (
	"fmt"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"
)

// EntityDiedService handles the death of an entity: notifies others via a System entity and then removes the dead entity.
func EntityDiedService(checksum string, token string) error {
	// 1. Validate Token and get UserID
	userID, err := utils.GetUserIDFromJWT(token)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}

	// 2. Check for existing "System" entity for this user
	// We want a unique system entity per user to manage their simulation's system messages.
	// We'll use a specific convention for system checksum: system_<userid>
	systemChecksum := fmt.Sprintf("system_%s", userID)

	exists, err := services.IsExist(systemChecksum)
	if err != nil {
		return fmt.Errorf("error checking system entity existence: %w", err)
	}

	if !exists {
		// Create System entity
		// Name="System", Checksum="system_<userid>", Role="SYSTEM", Prompt="System Announcer"
		_, err := services.CreateEntity("System", "System Announcer", systemChecksum, userID, "SYSTEM")
		if err != nil {
			return fmt.Errorf("failed to create system entity: %w", err)
		}
		pkg.DisplayContext("Created System entity for user "+userID, pkg.Update)
	}

	// 3. Get the name of the dying entity
	deadEntityName, err := services.GetEntityNameByChecksum(checksum)
	if err != nil {
		// If entity doesn't exist, maybe it's already dead?
		return fmt.Errorf("failed to get dying entity name: %w", err)
	}

	// 4. Broadcast message: "Entity <Name> has died." using System entity
	message := fmt.Sprintf("Entity %s has died.", deadEntityName)

	systemID, err := services.GetEntityIdByChecksum(systemChecksum)
	if err != nil {
		return fmt.Errorf("failed to get system entity ID: %w", err)
	}

	// We use the broadcast logic but from the System entity
	_, err = services.BroadcastMessage(intUserID, systemID, message)
	if err != nil {
		return fmt.Errorf("failed to broadcast death message: %w", err)
	}

	// // 5. Remove the dead entity
	// // Note: Messages sent BY the dead entity will be deleted due to CASCADE,
	// // but the death notification was sent by System, so it remains.
	// err = services.DropEntityByChecksum(checksum)
	// if err != nil {
	// 	return fmt.Errorf("failed to delete entity: %w", err)
	// }

	pkg.DisplayContext(fmt.Sprintf("Entity %s (%s) died and was removed.", deadEntityName, checksum), pkg.Update)

	return nil
}
