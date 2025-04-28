package sharedServices

import (
	"errors"
	"github.com/fatih/color"
	"my-api/internal/utils"
	"my-api/pkg"
)

// DisconnectService verifies a JWT token and disconnects the associated user if valid.
func DisconnectService(token string) error {
	result, err := utils.VerifyJWT(token)
	if err != nil {
		return errors.New("error during disconnection: " + err.Error())
	}

	if result != nil {
		color.Green("✅ Valid token. Disconnecting user ID: %v", result.UserID)
		pkg.DeleteToken(result.UserID)
	} else {
		color.Yellow("⚠️ Token is nil after verification")
	}

	return err
}
