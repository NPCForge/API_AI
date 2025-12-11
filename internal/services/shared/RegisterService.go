package sharedServices

import (
	"errors"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"
)

// RegisterService registers a new user by creating an account, generating a JWT, and storing the token.
func RegisterService(password string, identifier string, gamePrompt string) (string, string, error) {
	// Check if user already exists
	existingUserID, err := services.GetUserIdByName(identifier)
	if err == nil && existingUserID != -1 {
		// User exists, delete them
		err := services.DropUser(existingUserID)
		if err != nil {
			return "", "", errors.New("error deleting existing user: " + err.Error())
		}
	}

	id, err := services.Register(password, identifier, gamePrompt)
	if err != nil {
		pkg.DisplayContext(err.Error(), pkg.Error)
		return "", "", errors.New("error while registering to the database")
	}

	key, err := utils.GenerateJWT(strconv.Itoa(id))
	if err != nil {
		return "", "", errors.New("error while generating JWT")
	}

	pkg.SetToken(strconv.Itoa(id), key)

	return key, strconv.Itoa(id), nil
}
