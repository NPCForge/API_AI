package pkg

import (
	"log"
	"sync"
	"time"
)

var (
	tokenStore = make(map[string]string) // Clé: ID de l'utilisateur, Valeur: token
	mu         sync.Mutex
)

func SetToken(userID, token string) {
	log.Println("User", userID, "connected.")
	mu.Lock()
	defer mu.Unlock()
	tokenStore[userID] = token

	// Supprimer le token après une durée donnée (ex: 30 min)
	go func() {
		time.Sleep(30 * time.Minute)
		mu.Lock()
		delete(tokenStore, userID)
		mu.Unlock()
	}()
}

func GetToken(userID string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	token, exists := tokenStore[userID]
	return token, exists
}

func GetPopulation() map[string]string {
	mu.Lock()
	defer mu.Unlock()
	return tokenStore
}

func GetUserID(token string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()

	// Parcours du store pour trouver le token
	for userID, storedToken := range tokenStore {
		if storedToken == token {
			return userID, true // Renvoie l'ID de l'utilisateur si trouvé
		}
	}
	return "", false // Renvoie une chaîne vide si le token n'est pas trouvé
}

// Supprime le token d'un utilisateur
func DeleteToken(userID string) bool {
	mu.Lock()
	defer mu.Unlock()
	_, exists := tokenStore[userID]
	if exists {
		delete(tokenStore, userID)
	}
	return exists
}

// Met à jour le token d'un utilisateur (exemple basique qui régénère un token)
func UpdateToken(userID, newToken string) bool {
	mu.Lock()
	defer mu.Unlock()
	_, exists := tokenStore[userID]
	if exists {
		tokenStore[userID] = newToken
	}
	return exists
}

func IsValidToken(token string) bool {
	mu.Lock()
	defer mu.Unlock()

	for _, storedToken := range tokenStore {
		if storedToken == token {
			return true // Le token est valide
		}
	}
	return false
}

func ClearTokenStore() {
	mu.Lock()
	defer mu.Unlock()
	tokenStore = make(map[string]string)
}
