package service

import (
	"fmt"
	"my-api/internal/models/shared"
	"my-api/internal/services"
	"my-api/internal/utils"
)

func GetEntitiesService(self string) ([]sharedModel.Entity, error) {
	id, err := utils.GetUserIDFromJWT(self)

	if err != nil {
		fmt.Printf("GetEntitiesService: %s, token = %s\n", err, self)
		return nil, err
	}

	ids, checksums, err := services.GetEntities(id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var entities []sharedModel.Entity
	for i := range ids {
		entities = append(entities, sharedModel.Entity{
			Id:       ids[i],
			Checksum: checksums[i],
		})
	}

	return entities, nil
}
