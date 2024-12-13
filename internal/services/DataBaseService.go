package services

import (
	"database/sql"
	"fmt"
	"my-api/config"
	"my-api/internal/models"
)

// GetIDFromDB récupère l'ID correspondant à un token donné
func GetIDFromDB(token string) (int, error) {
	db := config.GetDB() // Récupère la connexion à la base de données

	var id int
	query := `SELECT id FROM entity WHERE token = $1`
	err := db.QueryRow(query, token).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil // Aucun enregistrement trouvé
	} else if err != nil {
		return 0, fmt.Errorf("erreur lors de la récupération de l'ID : %w", err)
	}

	return id, nil
}

// DropUser supprime un utilisateur en fonction de son token
func DropUser(token string) (string, error) {
	db := config.GetDB()

	// Supprimer directement par le token
	query := `DELETE FROM entity WHERE token = $1`
	result, err := db.Exec(query, token)

	if err != nil {
		return "", fmt.Errorf("erreur lors de la suppression de l'utilisateur : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("erreur lors de la vérification des lignes supprimées : %w", err)
	}

	if rowsAffected == 0 {
		return "", fmt.Errorf("aucun utilisateur avec ce token trouvé")
	}

	return "success", nil
}

// IsExist vérifie si un token existe dans la base de données
func IsExist(token string) (bool, error) {
	db := config.GetDB()

	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM entity WHERE token = $1)`
	err := db.QueryRow(query, token).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("erreur lors de la vérification de l'existence : %w", err)
	}

	return exists, nil
}

// Register insère une nouvelle entité dans la base de données
func Register(token string, entity models.RegisterRequest) (int64, error) {
	db := config.GetDB()

	query := `INSERT INTO entity (name, token, prompt, created) VALUES ($1, $2, $3, CURRENT_DATE) RETURNING id`

	var id int64
	err := db.QueryRow(query, entity.Name, token, entity.Prompt).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Error while registering entity : %w", err)
	}

	return id, nil
}
