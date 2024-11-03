package services

import (
	"encoding/json"
	"fmt"
	"my-api/config"
	"my-api/internal/models"
)

func GetIDFromDB(UserId string) (string, error) {
	client := config.GetClient() // Récupère le client configuré

	// Effectue une requête pour vérifier si un enregistrement avec ce token existe
	resp, err := client.R().
		SetQueryParam("filter", fmt.Sprintf(`token="%s"`, UserId)).
		Get("http://localhost:8090/api/collections/Entity/records")

	if err != nil {
		return "", fmt.Errorf("erreur lors de la requête : %w", err)
	}

	var response models.ViewReponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return "", fmt.Errorf("erreur lors du déchiffrage de la réponse : %w", err)
	}

	// Vérifie s'il y a des éléments
	if response.TotalItems > 0 && len(response.Items) > 0 {
		// Récupérer l'ID du premier élément trouvé
		id := response.Items[0].ID
		return id, nil
	} else {
		return "", nil // Aucune entrée avec ce token
	}
}

func DropUser(UserId string) (string, error) {
	client := config.GetClient()

	// Récupérer l'ID en base de données
	id, err := GetIDFromDB(UserId)
	if err != nil || id == "" {
		return "", fmt.Errorf("erreur lors de la récupération de l'ID : %w", err)
	}

	// Supprimer l'enregistrement directement par son ID
	resp, err := client.R().
		Delete(fmt.Sprintf("http://localhost:8090/api/collections/Entity/records/%s", id))

	// Vérifier si la suppression a réussi
	if err != nil {
		return "", fmt.Errorf("erreur lors de la suppression de l'utilisateur : %w", err)
	}
	if resp.StatusCode() != 204 {
		return "", fmt.Errorf("échec de la suppression de l'utilisateur, statut : %d", resp.StatusCode())
	}
	return "success", nil
}

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
		return true, nil // Une entrée avec le token existe
	} else {
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
