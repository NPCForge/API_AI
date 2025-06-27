package services

import (
	"database/sql"
	"fmt"
	"my-api/config"
	"my-api/internal/models"
	sharedModel "my-api/internal/models/shared"
	"my-api/pkg"
	"strings"

	"github.com/lib/pq"
)

// GetIDByChecksum get entity id for a given checksum
func GetIDByChecksum(checksum string) (int, error) {
	db := config.GetDB()

	var id int
	query := `SELECT id FROM entities WHERE checksum = $1`
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
	query := `SELECT prompt FROM entities WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&prompt)

	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("error while getting prompt : %w", err)
	}

	return prompt, nil
}

func GetNameByChecksum(checksum string) (string, error) {
	db := config.GetDB()

	var name string
	query := `SELECT name FROM entities WHERE checksum = $1`
	err := db.QueryRow(query, checksum).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("error while getting name : %w", err)
	}
	return name, nil
}

func GetNameByID(id string) (string, error) {
	db := config.GetDB()

	var name string
	query := `SELECT name FROM entities WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&name)

	if err != nil {
		return "", fmt.Errorf("error while getting name : %w", err)
	}

	return name, nil
}

func ResetGame() error {
	db := config.GetDB()

	query := "DELETE FROM discussions"

	_, err := db.Exec(query)

	if err != nil {
		return fmt.Errorf("error while getting name : %w", err)
	}

	return nil
}

func formatNewMessages(rows *sql.Rows, selfChecksum string) ([]string, error) {
	var formattedMessages []string

	for rows.Next() {
		var senderUserID int
		var receiverChecksums pq.StringArray
		var senderChecksum, messageContent string

		err := rows.Scan(&senderUserID, &senderChecksum, &messageContent, &receiverChecksums)
		if err != nil {
			pkg.DisplayContext("Error after row scan: ", pkg.Error, err)
			return nil, err
		}

		for i, value := range receiverChecksums {
			if value == selfChecksum {
				receiverChecksums[i] = "You"
			}
		}

		formattedMessages = append(
			formattedMessages,
			fmt.Sprintf("[%s -> %s: \"%s\"]", senderChecksum, strings.Join(receiverChecksums, ", "), messageContent),
		)
	}

	if err := rows.Err(); err != nil {
		pkg.DisplayContext("Error after rows iteration: ", pkg.Error, err)
		return nil, err
	}

	return formattedMessages, nil
}

func markNewMessagesAsRead(receiverID int) error {
	db := config.GetDB()

	queryUpdate := `
WITH filtered_messages AS (
	SELECT messages.id
	FROM messages
	JOIN message_receivers ON messages.id = message_receivers.message_id
	WHERE message_receivers.receiver_entity_id = $1
	  AND message_receivers.is_new_message = TRUE
	ORDER BY messages.timestamp
	LIMIT 5
)
UPDATE message_receivers
SET is_new_message = FALSE
WHERE message_id IN (SELECT id FROM filtered_messages)
  AND receiver_entity_id = $1;
`

	_, err := db.Exec(queryUpdate, receiverID)

	if err != nil {
		pkg.DisplayContext("[markNewMessagesAsRead]: ", pkg.Error, err)
		return err
	}
	return nil
}

// Refacto ✅
func GetNewMessages(ReceiverEntityChecksum string) ([]string, error) {
	db := config.GetDB()
	receiverId, err := GetIDByChecksum(ReceiverEntityChecksum)

	if err != nil {
		return nil, err
	}

	query := `WITH filtered_messages AS (
		SELECT messages.id, messages.sender_entity_id, messages.message, messages.timestamp
		FROM messages
		JOIN message_receivers ON messages.id = message_receivers.message_id
		WHERE message_receivers.receiver_entity_id = $1
		  AND message_receivers.is_new_message = TRUE
	)
	SELECT
		filtered_messages.sender_entity_id,
		entities.checksum AS EntityChecksum,
		filtered_messages.message,
		ARRAY_AGG(receiver_entity.checksum) AS ReceiverChecksum
		FROM filtered_messages
		JOIN entities ON filtered_messages.sender_entity_id = entities.id
		JOIN message_receivers ON filtered_messages.id = message_receivers.message_id
		JOIN entities AS receiver_entity ON message_receivers.receiver_entity_id = receiver_entity.id
		GROUP BY filtered_messages.id, filtered_messages.sender_entity_id, entities.checksum,
		filtered_messages.message, filtered_messages.timestamp
		ORDER BY filtered_messages.timestamp
		LIMIT 5;
	`

	rows, err := db.Query(query, receiverId)
	if err != nil {
		pkg.DisplayContext("Error after GetNewMessages query:", pkg.Error, err)
		return nil, err
	}

	formatedMessages, err := formatNewMessages(rows, ReceiverEntityChecksum)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			pkg.DisplayContext("Error closing rows:", pkg.Error, err)
		}
	}(rows)

	err = markNewMessagesAsRead(receiverId)

	if err != nil {
		return nil, err
	}

	return formatedMessages, nil
}

func GetDiscussions(EntityChecksum string) ([]sharedModel.Message, error) {
	db := config.GetDB()

	limitRows := 10

	query := `SELECT m.sender_entity_id, m.message
    FROM messages m
    ORDER BY m.timestamp
    LIMIT $1
`
	rows, err := db.Query(query, limitRows)
	if err != nil {
		pkg.DisplayContext("Error after GetDiscussion query:", pkg.Error, err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			pkg.DisplayContext("Error after row close:", pkg.Error, err)
		}
	}(rows)

	var messages []sharedModel.Message

	for rows.Next() {
		var senderEntityID int
		var msg sharedModel.Message

		err := rows.Scan(&senderEntityID, &msg.Message)
		if err != nil {
			pkg.DisplayContext("Error after row scan:", pkg.Error, err)
			return nil, err
		}

		//default Everyone
		msg.ReceiverChecksums = []string{"Everyone"}

		SenderChecksum, err := GetEntityChecksumById(senderEntityID)

		if err != nil {
			return nil, err
		}

		if SenderChecksum == EntityChecksum {
			msg.SenderChecksum = "You"
		} else {
			msg.SenderChecksum = SenderChecksum
		}

		//for i, r := range receiverChecksums {
		//	if r == EntityChecksum {
		//		receiverChecksums[i] = "You"
		//	}
		//}

		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		pkg.DisplayContext("Error after row scan:", pkg.Error, err)
		return nil, err
	}

	return messages, nil
}

func GetDiscussionFromBy(from string, to string) ([]sharedModel.Message, error) {
	db := config.GetDB()

	limitRows := 10

	query := `WITH filtered_messages AS (
    SELECT m.id, m.sender_entity_id, m.message, m.timestamp
    FROM messages m
    JOIN message_receivers mr ON m.id = mr.message_id
    WHERE (m.sender_entity_id = $1 AND mr.receiver_entity_id = $2)
       OR (m.sender_entity_id = $2 AND mr.receiver_entity_id = $1)
    ORDER BY m.timestamp
    LIMIT $3
)
SELECT 
    fm.sender_entity_id,
    sender_entity.checksum AS SenderChecksum,
    fm.message,
    ARRAY_AGG(receiver_entity.checksum) AS ReceiverChecksum
FROM filtered_messages fm
JOIN entities sender_entity ON fm.sender_entity_id = sender_entity.id
JOIN message_receivers mr ON fm.id = mr.message_id
JOIN entities receiver_entity ON mr.receiver_entity_id = receiver_entity.id
GROUP BY fm.id, fm.sender_entity_id, sender_entity.checksum, fm.message, fm.timestamp
ORDER BY fm.timestamp;
`

	rows, err := db.Query(query, from, to, limitRows)
	if err != nil {
		pkg.DisplayContext("Error after GetDiscussion query:", pkg.Error, err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			pkg.DisplayContext("Error after row close:", pkg.Error, err)
		}
	}(rows)

	var messages []sharedModel.Message

	for rows.Next() {
		var senderEntityID int
		var msg sharedModel.Message
		var receiverChecksums pq.StringArray

		err := rows.Scan(&senderEntityID, &msg.SenderChecksum, &msg.Message, &receiverChecksums)
		if err != nil {
			pkg.DisplayContext("Error after row scan:", pkg.Error, err)
			return nil, err
		}
		//msg.ReceiverChecksums = receiverChecksums

		if fmt.Sprintf("%d", senderEntityID) == from {
			msg.SenderChecksum = "You"
		}

		for i, r := range receiverChecksums {
			if r == from {
				receiverChecksums[i] = "You"
			}
		}

		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		pkg.DisplayContext("Error after row scan:", pkg.Error, err)
		return nil, err
	}

	return messages, nil
}

func GetUserEntities(userID int) ([]int, error) {
	db := config.GetDB()

	query := `SELECT id FROM entities WHERE user_id = $1`

	rows, err := db.Query(query, userID)

	if err != nil {
		pkg.DisplayContext("Broadcast message, cannot get user entities:", pkg.Error, err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			pkg.DisplayContext("Error after row close:", pkg.Error, err)
		}
	}(rows)

	var entities []int

	for rows.Next() {
		var entityID int
		err := rows.Scan(&entityID)
		if err != nil {
			pkg.DisplayContext("Error after row scan:", pkg.Error, err)
			return nil, err
		}
		entities = append(entities, entityID)
	}
	if err := rows.Err(); err != nil {
		pkg.DisplayContext("Error after row scan:", pkg.Error, err)
		return nil, err
	}
	return entities, nil
}

func BroadcastMessage(userID int, senderId int, message string) (int64, error) {
	db := config.GetDB()
	receiverIDs, err := GetUserEntities(userID)

	if err != nil {
		return -1, err
	}

	query := `INSERT INTO messages (sender_entity_id, message) VALUES ($1, $2) RETURNING id`

	var id int64
	err = db.QueryRow(query, senderId, message).Scan(&id)

	if err != nil {
		return -1, err
	}

	//self := globalHelpers.IntContains(receiverIDs, senderId)
	//
	//if self != -1 {
	//	receiverIDs = globalHelpers.RemoveIntAtIndex(receiverIDs, self)
	//}

	for _, receiverID := range receiverIDs {
		query = `INSERT INTO message_receivers (message_id, receiver_entity_id) VALUES ($1, $2)`

		_, err = db.Exec(query, id, receiverID)

		if err != nil {
			return 0, fmt.Errorf("error while insert message : %w", err)
		}
	}

	return 0, nil
}

func NewMessage(senderId int, receiverId int, message string) (int64, error) {
	db := config.GetDB()

	query := `INSERT INTO messages (sender_entity_id, message) VALUES ($1, $2) RETURNING id`

	var id int64
	err := db.QueryRow(query, senderId, message).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error while insert message : %w", err)
	}

	query = `INSERT INTO message_receivers (message_id, receiver_entity_id) VALUES ($1, $2)`

	_, err = db.Exec(query, id, receiverId)

	if err != nil {
		return 0, fmt.Errorf("error while insert message : %w", err)
	}

	return id, nil
}

func GetEntityIDByName(name string) (string, error) {
	db := config.GetDB()

	var entity string
	query := `SELECT id FROM entities WHERE LOWER(name) = LOWER($1)`
	err := db.QueryRow(query, name).Scan(&entity)

	if err == sql.ErrNoRows {
		return "Cannot find entity", err
	}

	return entity, nil
}

// refacto ✅
func DropUser(id int) error {
	db := config.GetDB()

	query := `DELETE FROM users WHERE id = $1`
	result, err := db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("erreur lors de la suppression de l'utilisateur : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification des lignes supprimées : %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("aucun utilisateur avec cet id trouvé")
	}

	return nil
}

func DropDiscussions() error {
	db := config.GetDB()

	query := `TRUNCATE TABLE message_receivers, messages RESTART IDENTITY CASCADE;`

	_, err := db.Exec(query)

	if err != nil {
		return fmt.Errorf("error while drop discussions: %w", err)
	}
	return nil
}

func IsExist(checksum string) (bool, error) {
	db := config.GetDB()

	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM entities WHERE checksum = $1)`
	err := db.QueryRow(query, checksum).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("erreur lors de la vérification de l'existence : %w", err)
	}

	return exists, nil
}

func IsExistById(id string) (bool, error) {
	db := config.GetDB()

	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM entities WHERE id = $1)`
	err := db.QueryRow(query, id).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("erreur lors de la vérification de l'existence : %w", err)
	}

	return exists, nil
}

// === Refacto === ✅

func Register(password string, identifier string, gamePrompt string) (int, error) {
	db := config.GetDB()

	// Hasher le mot de passe
	pass, err := pkg.HashPassword(password)

	if err != nil {
		return -1, fmt.Errorf("error hashing password: %w", err)
	}

	query := `
		INSERT INTO users (name, password_hash, game_prompt)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int
	err = db.QueryRow(query, identifier, pass, gamePrompt).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("error while registering user: %w", err)
	}

	return id, nil
}

func CreateEntity(name string, prompt string, checksum string, id_owner string, role string) (int, error) {
	db := config.GetDB()

	query := `
		INSERT INTO entities (user_id, name, checksum, prompt, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int
	err := db.QueryRow(query, id_owner, name, checksum, prompt, role).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("error while registering user: %w", err)
	}

	return id, nil
}

func GetEntities(id_owner string) (ids []string, checksums []string, err error) {
	db := config.GetDB()

	query := `SELECT id, checksum FROM entities WHERE user_id = $1`

	rows, err := db.Query(query, id_owner)

	if err != nil {
		return nil, nil, fmt.Errorf("error while getting entities: %w", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var id string
		var checksum string

		if err := rows.Scan(&id, &checksum); err != nil {
			return nil, nil, fmt.Errorf("error while scanning row: %w", err)
		}

		ids = append(ids, id)
		checksums = append(checksums, checksum)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return ids, checksums, nil
}

func DropEntityByChecksum(checksum string) error {
	db := config.GetDB()

	query := `
		DELETE FROM entities WHERE checksum = $1
	`

	_, err := db.Exec(query, checksum)
	if err != nil {
		return fmt.Errorf("error while deleting entity: %w", err)
	}

	return nil
}

func GetPermissionById(id int) (int, error) {
	db := config.GetDB()

	var perm int
	query := `SELECT permission FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&perm)

	if err != nil {
		return -1, fmt.Errorf("erreur lors de la vérification de l'existence : %w", err)
	}

	return perm, nil
}

func GetUserIdByName(name string) (int, error) {
	db := config.GetDB()

	var perm int
	query := `SELECT id FROM users WHERE name = $1`
	err := db.QueryRow(query, name).Scan(&perm)

	if err != nil {
		return -1, fmt.Errorf("erreur lors de la vérification de l'existence : %w", err)
	}

	return perm, nil
}

func Connect(password string, identifier string) (int, error) {
	db := config.GetDB()

	var userId int
	var pass string

	query := `SELECT id, password_hash FROM users WHERE LOWER(name) = LOWER($1)`
	err := db.QueryRow(query, identifier).Scan(&userId, &pass)

	if err != nil {
		return -1, fmt.Errorf("error while connecting user: %w", err)
	}

	if !pkg.CheckPasswordHash(password, pass) {
		return -1, fmt.Errorf("error while connecting user")
	}

	return userId, nil
}

func GetEntitiesByUserID(userID string) ([]models.Entity, error) {
	db := config.GetDB()

	query := `
		SELECT id, user_id, name, checksum, prompt, created
		FROM entities
		WHERE user_id = $1
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying entities: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			pkg.DisplayContext("Error after row close: ", pkg.Error, err)
		}
	}(rows)

	var entities []models.Entity

	for rows.Next() {
		var e models.Entity
		if err := rows.Scan(&e.ID, &e.UserID, &e.Name, &e.Checksum, &e.Prompt, &e.Created); err != nil {
			return nil, fmt.Errorf("error scanning entity: %w", err)
		}
		entities = append(entities, e)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return entities, nil
}

func GetEntitiesOwnerByChecksum(checksum string) (int, error) {
	db := config.GetDB()

	query := `
		SELECT user_id
		FROM entities
		WHERE checksum = $1
	`

	var id int
	err := db.QueryRow(query, checksum).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("error while fetching entity owner by checksum: %s : %w", checksum, err)
	}

	return id, nil
}

func GetEntityNameByChecksum(checksum string) (string, error) {
	db := config.GetDB()

	query := `
		SELECT name
		FROM entities
		WHERE checksum = $1
	`

	var name string
	err := db.QueryRow(query, checksum).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("error while getting entity name from checksum = %s : %w", checksum, err)
	}

	return name, nil
}

func GetGamePromptByUserID(UserID string) (string, error) {
	db := config.GetDB()

	query := `SELECT game_prompt FROM users WHERE id = $1`

	var prompt string
	err := db.QueryRow(query, UserID).Scan(&prompt)
	if err != nil {
		return "", fmt.Errorf("error while fetching prompt from JWT: %w", err)
	}
	return prompt, nil
}

func GetEntityRoleByChecksum(checksum string) (string, error) {
	db := config.GetDB()

	query := `SELECT role FROM entities WHERE checksum = $1`

	var role string
	err := db.QueryRow(query, checksum).Scan(&role)

	if err != nil {
		return "", fmt.Errorf("error while fetching role from checksum = %s : %w", checksum, err)
	}

	return role, nil
}

func GetEntityNameByID(id int) (string, error) {
	db := config.GetDB()
	query := `
SELECT name
FROM entities
WHERE id = $1
`

	var name string
	err := db.QueryRow(query, id).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("error while querying entity: %w", err)
	}

	return name, nil
}

func GetEntityChecksumById(id int) (string, error) {
	db := config.GetDB()
	query := `
SELECT checksum
FROM entities
WHERE id = $1
`
	var checksum string
	err := db.QueryRow(query, id).Scan(&checksum)
	if err != nil {
		return "", fmt.Errorf("error while fetching entity checksum: %w", err)
	}
	return checksum, nil
}

func GetEntityIdByChecksum(checksum string) (int, error) {
	db := config.GetDB()

	query := `
		SELECT id
		FROM entities
		WHERE checksum = $1
	`

	var id int
	err := db.QueryRow(query, checksum).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("error while getting entity id by checksum : %s : %w", checksum, err)
	}

	return id, nil
}
