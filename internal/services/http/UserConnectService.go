package httpServices

import (
	"errors"
	httpModels "my-api/internal/models/http"
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
	"strconv"
)

func UserConnect(req httpModels.ConnectRequest) (string, error) {
	response, err := services.IsExist(req.Token)

	if err != nil {
		return "", errors.New("error searching in table")
	}

	if !response {
		return "", errors.New("account didn't exist")
	}

	id, err := services.GetIDFromDB(req.Token)

	var stringId = strconv.Itoa(id)

	if err != nil {
		return "", errors.New("Error searching in database")
	}

	pass, err := utils.GenerateJWT(stringId)

	if err != nil {
		return "", errors.New("error generating JWT")
	}

	pkg.SetToken(stringId, pass)

	return pass, nil
}
