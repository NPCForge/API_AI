package services

import (
	"errors"
	"my-api/pkg"
)

func Remove(token string) (string, error) {
	UserId, err := pkg.GetUserID(token)

	if !err {
		return "failed", errors.New("error getting userid")
	}

	exist, err_ := IsExist(UserId)

	if err_ != nil {
		return "failed", errors.New("error using DB")
	}

	if !exist {
		return "Success", nil
	}

	response, err_ := DropUser(UserId)

	if err_ != nil || response == "" {
		return "failed", errors.New("error droping DB")
	}

	return "success", nil
}
