package services

import (
	"database/sql"
	"fmt"
	"my-api/config"
	httpModels "my-api/internal/models/http"
	websocketModels "my-api/internal/models/websocket"
	"my-api/pkg"

	"github.com/lib/pq"
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

// DropAllUser for debug
func DropAllUser() (int64, error) {
	db := config.GetDB()

	query := `DELETE FROM users`
	result, err := db.Exec(query)
	if err != nil {
		return 0, fmt.Errorf("erreur lors de la suppression : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("erreur lors de la récupération du nombre de lignes supprimées : %w", err)
	}

	return rowsAffected, nil
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

func GetNewMessages(receiver string) ([]websocketModels.Message, error) {
	db := config.GetDB()
	receiverName, err := GetNameByID(receiver)

	if err != nil {
		return nil, err
	}

	query := `WITH filtered_messages AS (
    SELECT m.id, m.sender_user_id, m.message, m.timestamp
    FROM messages m
    JOIN message_receivers mr ON m.id = mr.message_id
    WHERE mr.receiver_user_id = $1
    AND mr.is_new_message = TRUE
)
SELECT 
    fm.sender_user_id,
    sender_entity.name AS SenderName,
    fm.message,
    ARRAY_AGG(receiver_entity.name) AS ReceiverNames
FROM filtered_messages fm
JOIN entity sender_entity ON fm.sender_user_id = sender_entity.id
JOIN message_receivers mr ON fm.id = mr.message_id
JOIN entity receiver_entity ON mr.receiver_user_id = receiver_entity.id
GROUP BY fm.id, fm.sender_user_id, sender_entity.name, fm.message, fm.timestamp
ORDER BY fm.timestamp
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
		var senderUserID, messageID int
		var receiverNames pq.StringArray
		var msg websocketModels.Message

		err := rows.Scan(&senderUserID, &msg.SenderName, &msg.Message, &receiverNames)
		if err != nil {
			println("Error after row scan:", err.Error())
			return nil, err
		}
		msg.ReceiverNames = receiverNames

		for i, r := range receiverNames {
			if r == receiverName {
				receiverNames[i] = "You"
			}
		}

		messages = append(messages, msg)
		messageIDs = append(messageIDs, messageID)
	}

	if err := rows.Err(); err != nil {
		println("Error after rows")
		return nil, err
	}

	//if len(messageIDs) > 0 {
	//	updateQuery := `UPDATE message_receivers SET is_new_message = FALSE WHERE message_id IN (` + placeholders(len(messageIDs)) + `) AND receiver_user_id = $1;`
	//	args := make([]interface{}, len(messageIDs)+1)
	//	for i, id := range messageIDs {
	//		args[i] = id
	//	}
	//	args[len(messageIDs)] = receiver
	//
	//	_, err = db.Exec(updateQuery, args...)
	//	if err != nil {
	//		println("Error updating messages:", err.Error())
	//		return nil, err
	//	}
	//}

	return messages, nil
}

func GetDiscussion(from string, to string) ([]websocketModels.Message, error) {
	db := config.GetDB()
	receiverName, err := GetNameByID(from)

	if err != nil {
		return nil, err
	}

	query := `WITH filtered_messages AS (
    SELECT m.id, m.sender_user_id, m.message, m.timestamp
    FROM messages m
    JOIN message_receivers mr ON m.id = mr.message_id
    WHERE m.sender_user_id = $1
    AND mr.receiver_user_id = $2
  	OR m.sender_user_id = $2
  	AND mr.receiver_user_id = $1
)
SELECT 
    fm.sender_user_id,
    sender_entity.name AS SenderName,
    fm.message,
    ARRAY_AGG(receiver_entity.name) AS ReceiverNames
FROM filtered_messages fm
JOIN entity sender_entity ON fm.sender_user_id = sender_entity.id
JOIN message_receivers mr ON fm.id = mr.message_id
JOIN entity receiver_entity ON mr.receiver_user_id = receiver_entity.id
GROUP BY fm.id, fm.sender_user_id, sender_entity.name, fm.message, fm.timestamp
ORDER BY fm.timestamp`

	rows, err := db.Query(query, from, to)
	if err != nil {
		println("Error after query")
		return nil, err
	}
	defer rows.Close()

	var messages []websocketModels.Message

	for rows.Next() {
		var senderUserID int
		var msg websocketModels.Message
		var receiverNames pq.StringArray

		err := rows.Scan(&senderUserID, &msg.SenderName, &msg.Message, &receiverNames)
		if err != nil {
			println("Error after row scan:", err.Error())
			return nil, err
		}
		msg.ReceiverNames = receiverNames

		if fmt.Sprintf("%d", senderUserID) == from {
			msg.SenderName = "You"
		}

		for i, r := range receiverNames {
			if r == receiverName {
				receiverNames[i] = "You"
			}
		}

		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		println("Error after rows")
		return nil, err
	}

	return messages, nil
}

func NewMessage(senderId int, receiverId int, message string) (int64, error) {
	db := config.GetDB()

	query := `INSERT INTO messages (sender_user_id, message) VALUES ($1, $2) RETURNING id`

	var id int64
	err := db.QueryRow(query, senderId, message).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error while insert message : %w", err)
	}

	query = `INSERT INTO message_receivers (message_id, receiver_user_id) VALUES ($1, $2)`

	_, err = db.Exec(query, id, receiverId)

	if err != nil {
		return 0, fmt.Errorf("error while insert message : %w", err)
	}

	return id, nil
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

func RegisterRefacto(password string, identifier string) (int, error) {
	db := config.GetDB()

	// Hasher le mot de passe
	pass, err := pkg.HashPassword(password)
	if err != nil {
		return -1, fmt.Errorf("error hashing password: %w", err)
	}

	query := `
		INSERT INTO users (name, password_hash, created)
		VALUES ($1, $2, CURRENT_DATE)
		RETURNING id
	`

	var id int
	err = db.QueryRow(query, identifier, pass).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("error while registering entity: %w", err)
	}

	return id, nil
}
