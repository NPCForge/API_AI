package services

import (
	"errors"
	"my-api/internal/models"
	"my-api/pkg"
)

func UserConnect(req models.ConnectRequest) (string, error) {
	response, err := IsExist(req.Token)

	if err != nil {
		return "", errors.New("error searching in table")
	}

	if !response {
		return "", errors.New("account didn't exist")
	}

	pass, err := pkg.GenerateJWT(req.Token)

	pkg.SetToken(req.Token, pass)

	if err != nil {
		return "", errors.New("error generating JWT")
	}

	return pass, nil
}
