package config

import (
	"fmt"
	. "my-api/pkg"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var displayDebug bool

func setDefaultVariable() {
	displayDebug = false
	if val := GetEnvVariable("DISPLAY_DEBUG"); strings.ToLower(val) == "true" {
		displayDebug = true
	}
}

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
