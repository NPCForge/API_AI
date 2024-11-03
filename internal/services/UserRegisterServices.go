package services

import (
	"errors"
	"my-api/internal/models"
	"my-api/pkg"
)

func SaveInDatabase(token string, entity models.RegisterRequest) (string, error) {
	response, err := IsExist(token)

	if err != nil {
		return "", errors.New("error searching in table")
	}

	if response {
		return "", errors.New("error entry already exist in database")
	}

	response, err = Register(token, entity)

	if err != nil || !response {
		return "", errors.New("error creating entry")
	}

	return token, nil
}

// Generation d'un token, et enregistrement de l'entité dans la base de donnée
func RegisterService(entity models.RegisterRequest) (string, error) {
	pass, err := pkg.GenerateJWT(entity.Name)

	if err != nil {
		return "", errors.New("error generating JWT")
	}

	pass, err = SaveInDatabase(pass, entity)

	if err != nil {
		return "", errors.New("error saving in database")
	}

	return pass, nil
}
