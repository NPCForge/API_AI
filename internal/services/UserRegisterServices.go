package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"my-api/config"
	"my-api/internal/models"
	"my-api/pkg"
)

func IsExist(token string) (bool, error) {
	client := config.GetClient() // Récupère le client configuré

	// Effectue une requête pour vérifier si un enregistrement avec ce token existe
	resp, err := client.R().
		SetQueryParam("filter", fmt.Sprintf(`token="%s"`, token)).
		Get("http://localhost:8090/api/collections/Entity/records")

	if err != nil {
		return false, fmt.Errorf("erreur lors de la requête : %w", err)
	}

	// Dé-sérialise la réponse JSON
	var response models.ViewReponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return false, fmt.Errorf("erreur lors du déchiffrage de la réponse : %w", err)
	}

	// Vérifie s'il y a des éléments
	if response.TotalItems > 0 {
		fmt.Println("Entrée trouvée :", response.Items)
		return true, nil // Une entrée avec le token existe
	} else {
		fmt.Println("Aucune entrée trouvée avec ce token.")
		return false, nil // Aucune entrée avec ce token
	}
}

func Register(token string, entity models.RegisterRequest) (bool, error) {
	client := config.GetClient()

	var body models.BodyEntity = models.BodyEntity{
		Token:  token,
		Prompt: entity.Prompt,
		Name:   entity.Name,
	}

	resp, err := client.R().
		SetBody(body).
		Post("http://localhost:8090/api/collections/Entity/records")

	if err != nil {
		return false, fmt.Errorf("erreur lors de la requête : %w", err)
	}

	if resp.StatusCode() != 200 {
		return false, fmt.Errorf("erreur lors de la requête status !200")
	}
	return true, nil
}

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
