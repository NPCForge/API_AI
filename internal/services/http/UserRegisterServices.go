package http

import (
	"errors"
	"my-api/internal/models/http"
	"my-api/internal/services"
	"my-api/pkg"
	"strconv"
)

func SaveInDatabase(entity http.RegisterRequest) (int64, error) {
	response, err := services.IsExist(entity.Token)

	if err != nil {
		return -1, errors.New("error searching in table")
	}

	if response {
		return -1, errors.New("error entry already exist in database")
	}

	id, err := services.Register(entity.Token, entity)

	if err != nil {
		return -1, errors.New("error while registering")
	}

	return id, nil
}

// Generation d'un token, et enregistrement de l'entité dans la base de donnée
func RegisterService(entity http.RegisterRequest) (string, error) {
	id, err := SaveInDatabase(entity)

	if err != nil {
		return "", errors.New("error saving in database")
	}

	pass, err := pkg.GenerateJWT(strconv.FormatInt(id, 10))

	if err != nil {
		return "", errors.New("error generating JWT")
	}

	return pass, nil
}
