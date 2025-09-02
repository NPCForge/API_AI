package sharedServices

import (
	"my-api/internal/services"
	"my-api/internal/utils"
	"my-api/pkg"
)

// CreateEntityService creates a new entity linked to a user identified by a JWT token.
func CreateEntityService(name string, prompt string, checksum string, Token string, role string) (int, error) {
	id, err := utils.GetUserIDFromJWT(Token)
	if err != nil {
		pkg.DisplayContext(err.Error(), pkg.Error)
		return -1, err
	}

	res, err := services.CreateEntity(name, prompt, checksum, id, role)
	if err != nil {
		pkg.DisplayContext(err.Error(), pkg.Error)
		return -1, err
	}

	return res, nil
}
