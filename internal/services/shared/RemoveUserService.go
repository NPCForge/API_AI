package sharedServices

import (
	"fmt"
	"my-api/internal/services"
	"my-api/internal/utils"
	"strconv"
)

// current_user: JWT token
// delete_user: username (to be deleted)
func RemoveUserService(current_user string, delete_user string) error {
	// Extract user ID from the JWT token
	currentIDStr, err := utils.GetUserIDFromJWT(current_user)
	if err != nil {
		return err
	}

	currentID, err := strconv.Atoi(currentIDStr)
	if err != nil {
		return err
	}

	// Get the current user's permission level
	perm, err := services.GetPermissionById(currentID)
	if err != nil {
		return err
	}

	// Get the ID of the user to be deleted
	deleteUserId := -1

	if delete_user != "" {
		deleteUserId, err = services.GetUserIdByName(delete_user)
		if err != nil {
			return err
		}
	} else {
		deleteUserId = currentID
	}

	isSelf := deleteUserId == currentID

	// Standard users can only delete their own account
	if perm == 0 && isSelf {
		err = services.DropUser(currentID)
		if err != nil {
			return err
		}

		err := DisconnectService(current_user)
		if err != nil {
			return err
		}

		return nil
	} else if perm == 0 {
		return fmt.Errorf("users can only delete their own account")
	}

	// Admins can delete any user
	fmt.Println("Deletion authorized for user ID", deleteUserId)

	err = services.DropUser(deleteUserId)

	if err != nil {
		return err
	}

	err = DisconnectService(delete_user)
	if err != nil {
		return err
	}

	return nil
}
