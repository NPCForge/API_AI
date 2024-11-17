package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	dbClient *sql.DB
	once     sync.Once
)

func InitDB() {
	once.Do(func() {
		// Configuration de la connexion
		connStr := "user=" + GetEnvVariable("DB_USER")
		connStr += " password=" + GetEnvVariable("DB_PASSWORD")
		connStr += " dbname=" + GetEnvVariable("DB_NAME")
		connStr += " host=" + GetEnvVariable("DB_HOST")
		connStr += " port=" + GetEnvVariable("DB_PORT")
		connStr += " sslmode=disable"

		// Se connecter à la base de données
		var err error
		dbClient, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Erreur lors de l'ouverture de la connexion à PostgreSQL : %v", err)
		}

		// Tester la connexion
		err = dbClient.Ping()
		if err != nil {
			log.Fatalf("Erreur de connexion à PostgreSQL : %v", err)
		}

		fmt.Println("Connexion à PostgreSQL réussie !")
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
