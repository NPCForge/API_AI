package services

import (
	"database/sql"
	"fmt"
	"strings"

	"my-api/config"
	httpModels "my-api/internal/models/http"
	websocketModels "my-api/internal/models/websocket"
)

// GetIDFromDB récupère l'ID correspondant à un checksum donné
func GetIDFromDB(checksum string) (int, error) {
	db := config.GetDB()

	var id int
	query := `SELECT id FROM entity WHERE checksum = $1`
	err := db.QueryRow(query, checksum).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil
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

func GetNameByID(id string) (string, error) {
	db := config.GetDB()

	var name string
	query := `SELECT name FROM entity WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&name)

	if err != nil {
		return "", fmt.Errorf("error while getting name : %w", err)
	}

	return name, nil
}

func placeholders(n int) string {
	placeholders := make([]string, n)
	for i := range placeholders {
		placeholders[i] = "$" + fmt.Sprintf("%d", i+1)
	}
	return strings.Join(placeholders, ", ")
}

func GetNewMessages(receiver string) ([]websocketModels.Message, error) {
	db := config.GetDB()

	query := `
	SELECT d.id, d.sender_user_id, d.receiver_user_id, e1.name AS sender_name, e2.name AS receiver_name, d.message, d.is_new_message
	FROM discussions d
	JOIN entity e1 ON d.sender_user_id = e1.id
	JOIN entity e2 ON d.receiver_user_id = e2.id
	WHERE d.receiver_user_id = $1 AND d.is_new_message = TRUE
	ORDER BY d.timestamp
	LIMIT 5;
	`

	rows, err := db.Query(query, receiver)
	if err != nil {
		println("Error after query:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var messages []websocketModels.Message
	var messageIDs []int

	for rows.Next() {
		var msg websocketModels.Message
		var senderUserID, receiverUserID, messageID int

		err := rows.Scan(&messageID, &senderUserID, &receiverUserID, &msg.SenderName, &msg.ReceiverName, &msg.Message, &msg.IsNewMessage)
		if err != nil {
			println("Error after row scan:", err.Error())
			return nil, err
		}

		if fmt.Sprintf("%d", receiverUserID) == receiver {
			msg.ReceiverName = "You"
		}

		messages = append(messages, msg)
		messageIDs = append(messageIDs, messageID)
	}

	if err := rows.Err(); err != nil {
		println("Error after rows")
		return nil, err
	}

	if len(messageIDs) > 0 {
		updateQuery := `UPDATE discussions SET is_new_message = FALSE WHERE id IN (` + placeholders(len(messageIDs)) + `);`
		args := make([]interface{}, len(messageIDs))
		for i, id := range messageIDs {
			args[i] = id
		}

		_, err = db.Exec(updateQuery, args...)
		if err != nil {
			println("Error updating messages:", err.Error())
			return nil, err
		}
	}

	return messages, nil
}

func GetDiscussion(from string, to string) ([]websocketModels.Message, error) {
	db := config.GetDB()

	query := `
	SELECT d.sender_user_id, d.receiver_user_id, e1.name AS sender_name, e2.name AS receiver_name, d.message, d.is_new_message
	FROM discussions d
	JOIN entity e1 ON d.sender_user_id = e1.id
	JOIN entity e2 ON d.receiver_user_id = e2.id
	WHERE (d.sender_user_id = $1 AND d.receiver_user_id = $2)
	   OR (d.sender_user_id = $2 AND d.receiver_user_id = $1)
	ORDER BY d.timestamp;
	`

	rows, err := db.Query(query, from, to)
	if err != nil {
		println("Error after query")
		return nil, err
	}
	defer rows.Close()

	var messages []websocketModels.Message

	for rows.Next() {
		var msg websocketModels.Message
		var senderUserID, receiverUserID int

		err := rows.Scan(&senderUserID, &receiverUserID, &msg.SenderName, &msg.ReceiverName, &msg.Message, &msg.IsNewMessage)
		if err != nil {
			println("Error after row scan:", err.Error())
			return nil, err
		}

		if fmt.Sprintf("%d", senderUserID) == from {
			msg.SenderName = "You"
		}
		if fmt.Sprintf("%d", receiverUserID) == from {
			msg.ReceiverName = "You"
		}

		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		println("Error after rows")
		return nil, err
	}

	return messages, nil
}

func GetEntityByName(name string) (string, error) {
	db := config.GetDB()

	var entity string
	query := `SELECT id FROM entity WHERE LOWER(name) = LOWER($1)`
	err := db.QueryRow(query, name).Scan(&entity)

	if err == sql.ErrNoRows {
		return "Cannot find entity", err
	}

	return entity, nil
}

func DropUser(id string) (string, error) {
	db := config.GetDB()

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
