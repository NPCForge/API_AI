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
	// We blindly attempt to delete the user by name to ensure we can recreate it.
	// This avoids race conditions or issues where Select fails but Insert collision occurs.
	err := services.DropUserByName(identifier)
	if err != nil {
		pkg.DisplayContext("Error attempting to delete user by name: "+identifier, pkg.Error, err)
		// We proceed anyway, as the error might be benign, or the register might fail with a descriptive error.
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
