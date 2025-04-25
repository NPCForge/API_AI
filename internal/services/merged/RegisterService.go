package service

import (
	"errors"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"
)

func RegisterService(password string, identifier string) (string, string, error) {
	// pkg.DisplayContext("RegisterService", pkg.Debug)
	id, err := services.RegisterRefacto(password, identifier)

	if err != nil {
		return "", "", errors.New("error while registering, DB")
	}

	key, err := utils.GenerateJWT(strconv.Itoa(id))

	if err != nil {
		return "", "", errors.New("error while registering, JWT")
	}

	pkg.SetToken(strconv.Itoa(id), key)

	return key, strconv.Itoa(id), nil
}
