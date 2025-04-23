package config

import (
	"fmt"
	. "my-api/pkg"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVariable load environment variable from .env.local
func GetEnvVariable(c string) string {
	if err := godotenv.Load(".env.local"); err != nil {
		DisplayContext("Error while loading .env.local", Error, true)
	}

	variable := os.Getenv(c)
	if variable == "" {
		DisplayContext(fmt.Sprintf("%s undefined in .env.local", c), Error, true)
	}

	return variable
}
