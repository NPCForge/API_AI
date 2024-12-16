package http

import (
	"errors"
	"my-api/internal/services"
)

func MakeDecisionService(req string) (string, error) {
	res, err := services.GptSimpleRequest(req)

	if err != nil {
		return "", errors.New("error requesting Chatgpt")
	}

	return res, nil
}
