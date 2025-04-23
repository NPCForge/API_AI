package service

import (
	"errors"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"
	//"my-api/internal/utils"
	//"my-api/pkg"
	//"strconv"
)

func ConnectService(password string, identifier string) (string, error) {
	id, err := services.ConnectRefacto(password, identifier)

	if err != nil {
		return "", errors.New("error connecting service")
	}

	key, err := utils.GenerateJWT(strconv.Itoa(id))

	if err != nil {
		return "", errors.New("error while registering, JWT")
	}

	pkg.SetToken(strconv.Itoa(id), key)

	return key, nil
}
