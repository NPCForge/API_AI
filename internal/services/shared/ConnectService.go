package sharedServices

import (
	"errors"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"
)

// ConnectService authenticates a user by password and identifier, generates a JWT token, and stores it.
func ConnectService(password string, identifier string) (string, string, error) {
	id, err := services.Connect(password, identifier)
	if err != nil {
		return "", "", errors.New("error connecting service")
	}

	key, err := utils.GenerateJWT(strconv.Itoa(id))
	if err != nil {
		return "", "", errors.New("error while generating JWT")
	}

	pkg.SetToken(strconv.Itoa(id), key)

	return key, strconv.Itoa(id), nil
}
