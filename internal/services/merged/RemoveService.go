package service

import (
	"fmt"
	"my-api/internal/services"
	"my-api/internal/utils"
	"strconv"
)

// current_user: JWT token
// delete_user: username (to be deleted)
func RemoveService(current_user string, delete_user string) (bool, error) {
	// Extract user ID from the JWT token
	currentIDStr, err := utils.GetUserIDFromJWT(current_user)
	if err != nil {
		return false, err
	}

	currentID, err := strconv.Atoi(currentIDStr)
	if err != nil {
		return false, err
	}

	// Get the current user's permission level
	perm, err := services.GetPermissionByIdRefacto(currentID)
	if err != nil {
		return false, err
	}

	// Get the ID of the user to be deleted
	deletedUserID, err := services.GetUserIdByNameRefacto(delete_user)
	if err != nil {
		return false, err
	}

	isSelf := deletedUserID == currentID

	// Standard users can only delete their own account
	if perm == 0 {
		if !isSelf {
			return false, fmt.Errorf("standard users can only delete their own account")
		}

		err = services.DropUser(currentID)
		if err != nil {
			return false, err
		}

		return isSelf, nil // the user should be disconnected
	}

	// Admins can delete any user
	fmt.Println("Deletion authorized for user ID", deletedUserID)

	err = services.DropUser(deletedUserID)
	if err != nil {
		return false, err
	}

	return isSelf, nil
}
