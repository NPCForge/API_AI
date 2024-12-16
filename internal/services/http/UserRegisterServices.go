package httpServices

import (
	"errors"
	"my-api/internal/models/http"
	"my-api/internal/services"
	"my-api/pkg"
	"strconv"
)

func SaveInDatabase(entity httpModels.RegisterRequest) (int64, error) {
	response, err := services.IsExist(entity.Token)

	if err != nil {
		return -1, errors.New("error searching in table")
	}

	if response {
		return -1, errors.New("entry already exists in the database")
	}

	id, err := services.Register(entity.Checksum, entity)

	if err != nil {
		return -1, errors.New("error while registering")
	}

	return id, nil
}

// Token generation and entity registration in the database
func RegisterService(entity httpModels.RegisterRequest) (string, error) {
	id, err := SaveInDatabase(entity)

	if err != nil {
		return "", errors.New("error saving in database")
	}

	pass, err := pkg.GenerateJWT(strconv.FormatInt(id, 10))

	if err != nil {
		return "", errors.New("error generating JWT")
	}

	pkg.SetToken(strconv.FormatInt(id, 10), pass)

	return pass, nil
}
