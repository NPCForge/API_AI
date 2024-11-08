package services

import "errors"

func MakeDecisionService(req string) (string, error) {
	res, err := GptSimpleRequest(req)

	if err != nil {
		return "", errors.New("error requesting Chatgpt")
	}

	return res, nil
}
