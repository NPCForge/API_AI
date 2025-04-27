package service

import (
	"my-api/internal/services"
	"my-api/internal/utils"
)

func CreateEntityService(name string, prompt string, checksum string, self string) (int, error) {
	id, err := utils.GetUserIDFromJWT(self)

	if err != nil {
		return -1, err
	}

	res, err := services.CreateEntity(name, prompt, checksum, id)

	if err != nil {
		return -1, err
	}
	return res, nil
}
