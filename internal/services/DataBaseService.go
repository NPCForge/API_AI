package services

import (
	"database/sql"
	"fmt"
	"my-api/config"
	httpModels "my-api/internal/models/http"
	websocketModels "my-api/internal/models/websocket"
)

// GetIDFromDB récupère l'ID correspondant à un checksum donné
func GetIDFromDB(checksum string) (int, error) {
	db := config.GetDB() // Récupère la connexion à la base de données

	var id int
	query := `SELECT id FROM entity WHERE checksum = $1`
	err := db.QueryRow(query, checksum).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil // Aucun enregistrement trouvé
	} else if err != nil {
		return 0, fmt.Errorf("erreur lors de la récupération de l'id : %w", err)
	}

	return id, nil
}

func GetPromptByID(id string) (string, error) {
	db := config.GetDB()

	var prompt string
	query := `SELECT prompt FROM entity WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&prompt)
	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("error while getting prompt : %w", err)
	}

	return prompt, nil
}

// DropUser supprime un utilisateur en fonction de son id
func DropUser(id string) (string, error) {
	db := config.GetDB()

	// Supprimer directement par le checksum
	query := `DELETE FROM entity WHERE id = $1`

	result, err := db.Exec(query, id)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la suppression de l'utilisateur : %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return "", fmt.Errorf("erreur lors de la vérification des lignes supprimées : %w", err)
	}

	if rowsAffected == 0 {
		return "", fmt.Errorf("aucun utilisateur avec cet id trouvé")
	}

	return "success", nil
}

// IsExist vérifie si un checksum existe dans la base de données
func IsExist(checksum string) (bool, error) {
	db := config.GetDB()

	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM entity WHERE checksum = $1)`
	err := db.QueryRow(query, checksum).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("erreur lors de la vérification de l'existence : %w", err)
	}

	return exists, nil
}

// IsExistById vérifie si un id existe dans la base de données
func IsExistById(id string) (bool, error) {
	db := config.GetDB()

	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM entity WHERE id = $1)`
	err := db.QueryRow(query, id).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("erreur lors de la vérification de l'existence : %w", err)
	}

	return exists, nil
}

// Register insère une nouvelle entité dans la base de données
func Register(checksum string, entity httpModels.RegisterRequest) (int64, error) {
	db := config.GetDB()

	query := `INSERT INTO entity (name, checksum, prompt, created) VALUES ($1, $2, $3, CURRENT_DATE) RETURNING id`

	var id int64
	err := db.QueryRow(query, entity.Name, checksum, entity.Prompt).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error while registering entity : %w", err)
	}

	return id, nil
}

func RegisterWebsocket(checksum string, entity websocketModels.RegisterRequest) (int64, error) {
	db := config.GetDB()

	query := `INSERT INTO entity (name, checksum, prompt, created) VALUES ($1, $2, $3, CURRENT_DATE) RETURNING id`

	var id int64
	err := db.QueryRow(query, entity.Name, checksum, entity.Prompt).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error while registering entity : %w", err)
	}

	return id, nil
}

func NewMessage(senderId int, receiverId string, message string) (int64, error) {
	db := config.GetDB()

	query := `INSERT INTO discussions (sender_user_id, receiver_user_id, message) VALUES ($1, $2, $3) RETURNING id`

	var id int64
	err := db.QueryRow(query, senderId, receiverId, message).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error while insert message : %w", err)
	}

	return id, nil
}
