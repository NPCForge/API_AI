package sharedServices

import (
	"errors"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"
)

// RegisterService registers a new user by creating an account, generating a JWT, and storing the token.
func RegisterService(password string, identifier string) (string, string, error) {
	id, err := services.Register(password, identifier)
	if err != nil {
		return "", "", errors.New("error while registering to the database")
	}

	key, err := utils.GenerateJWT(strconv.Itoa(id))
	if err != nil {
		return "", "", errors.New("error while generating JWT")
	}

	pkg.SetToken(strconv.Itoa(id), key)

	return key, strconv.Itoa(id), nil
}
