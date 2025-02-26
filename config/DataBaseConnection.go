package config

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	dbClient *sql.DB
	once     sync.Once
)

func IsRunningInDocker() bool {
	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "/docker/") || strings.Contains(scanner.Text(), "/kubepods/") {
			return true
		}
	}
	return false
}

func InitDB() {
	once.Do(func() {
		// Configuration de la connexion
		connStr := "user=" + GetEnvVariable("POSTGRES_USER")
		connStr += " password=" + GetEnvVariable("POSTGRES_PASSWORD")
		connStr += " dbname=" + GetEnvVariable("POSTGRES_DB")

		if IsRunningInDocker() {
			connStr += " host=postgres"
		} else {
			connStr += " host=" + GetEnvVariable("POSTGRES_HOST")
		}
		connStr += " port=" + GetEnvVariable("POSTGRES_PORT")
		connStr += " sslmode=disable"

		var err error
		maxRetries := 4
		retryDelay := 10 * time.Second

		// Réessayer de se connecter plusieurs fois
		for retries := 0; retries < maxRetries; retries++ {
			// Tenter d'ouvrir la connexion
			dbClient, err = sql.Open("postgres", connStr)
			if err != nil {
				log.Printf("Erreur lors de l'ouverture de la connexion à PostgreSQL (tentative %d/%d) : %v", retries+1, maxRetries, err)
			} else {
				// Tester la connexion
				err = dbClient.Ping()
				if err == nil {
					log.Println("Connexion à PostgreSQL réussie !")
					return
				}

				log.Printf("Erreur de connexion à PostgreSQL (tentative %d/%d) : %v", retries+1, maxRetries, err)
			}

			// Fermer la connexion en cas d'erreur pour éviter les fuites
			if dbClient != nil {
				dbClient.Close()
			}

			// Attendre avant la prochaine tentative
			if retries < maxRetries-1 {
				log.Printf("Nouvelle tentative dans %v...", retryDelay)
				time.Sleep(retryDelay)
			}
		}

		// Si toutes les tentatives échouent
		log.Fatalf("Échec de connexion à PostgreSQL après %d tentatives : %v", maxRetries, err)
	})
}

// GetDB retourne le client PostgreSQL
func GetDB() *sql.DB {
	if dbClient == nil {
		InitDB()
	}
	return dbClient
}

// CloseDB ferme la connexion à PostgreSQL
func CloseDB() {
	if dbClient != nil {
		err := dbClient.Close()
		if err != nil {
			log.Printf("Erreur lors de la fermeture de la connexion à PostgreSQL : %v", err)
		} else {
			fmt.Println("Connexion à PostgreSQL fermée.")
		}
	}
}
