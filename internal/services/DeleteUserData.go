package services

import (
	"fmt"

	"my-api/config"
)

// DropEntitiesByUserID removes all entities (and related data) for a given user.
func DropEntitiesByUserID(userID int) error {
	db := config.GetDB()

	_, err := db.Exec("DELETE FROM entities WHERE user_id = $1", userID)
	if err != nil {
		return fmt.Errorf("error while deleting user entities: %w", err)
	}
	return nil
}