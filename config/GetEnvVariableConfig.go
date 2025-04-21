package config

import (
	"fmt"
	. "my-api/pkg"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVariable charge une variable d'environnement depuis .env.local
func GetEnvVariable(c string) string {
	// Charge le fichier .env.local
	if err := godotenv.Load(".env.local"); err != nil {
		DisplayContext("Erreur de chargement du fichier .env.local", Error, true)
	}

	// Récupère la variable d'environnement
	variable := os.Getenv(c)
	if variable == "" {
		DisplayContext(fmt.Sprintf("%s non définie dans .env.local", c), Error, true)
	}

	return variable
}
