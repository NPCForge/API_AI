package sharedServices

import (
	"errors"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"

	"github.com/fatih/color"
)

// DisconnectService verifies a JWT token and disconnects the associated user if valid.
func DisconnectService(token string) error {
	result, err := utils.VerifyJWT(token)
	if err != nil {
		return errors.New("error during disconnection: " + err.Error())
	}

	if result != nil {
		color.Green("✅ Valid token. Disconnecting user ID: %v", result.UserID)

		// Convert UserID from string to int
		userID, err := strconv.Atoi(result.UserID)
		if err != nil {
			return errors.New("invalid user ID format")
		}

		// Delete the user from the database
		// This will cascade delete their data (entities, messages, etc.)
		err = services.DropUser(userID)
		if err != nil {
			return errors.New("error deleting user: " + err.Error())
		}

		pkg.DeleteToken(result.UserID)
	} else {
		color.Yellow("⚠️ Token is nil after verification")
	}

	return err
}
