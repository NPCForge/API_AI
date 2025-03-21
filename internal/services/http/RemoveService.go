package httpServices

import (
	"errors"
	"my-api/internal/services"
	"my-api/pkg"
)

func Remove(token string) (string, error) {

	UserId, err := pkg.GetUserIDFromJWT(token)

	if err != nil {
		return "failed", errors.New("error getting userid")
	}

	exist, err_ := services.IsExistById(UserId)

	if err_ != nil {
		return "failed", errors.New("error using DB")
	}

	if !exist {
		return "Success", nil
	}

	response, err_ := services.DropUser(UserId)

	if err_ != nil || response == "" {
		return "failed", errors.New("error droping DB")
	}

	return "success", nil
}
