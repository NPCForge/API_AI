package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVariable charge une variable d'environnement depuis .env.local
func GetEnvVariable(c string) string {
	// Charge le fichier .env.local
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal("Erreur de chargement du fichier .env.local")
	}

	// Récupère la variable d'environnement
	variable := os.Getenv(c)
	if variable == "" {
		log.Fatal(fmt.Sprintf("%s non définie dans .env.local", c))
	}

	return variable
}
