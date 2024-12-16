package http

import (
	"errors"
	"my-api/internal/models/http"
	"my-api/internal/services"
	"my-api/pkg"
)

func UserConnect(req http.ConnectRequest) (string, error) {
	response, err := services.IsExist(req.Token)

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
