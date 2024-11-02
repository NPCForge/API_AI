package services

import (
	"errors"
	"fmt"
	"my-api/pkg"
)

func Disconnect(token string) (string, error) {
	userid, err := pkg.GetUserID(token)

	if !err {
		return "failed", errors.New("error getting userid")
	}

	res := pkg.DeleteToken(userid)

	if !res {
		fmt.Printf("Disconnect")
		return "failed", errors.New("error didn't exist")
	}

	return "success", nil
}
