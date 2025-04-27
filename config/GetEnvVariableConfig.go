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
	// Always load .env.local (it's safe)
	_ = godotenv.Load(".env.local")

	// Now read the environment variable
	variable := os.Getenv(c)
	if variable == "" {
		DisplayContext(fmt.Sprintf("%s undefined", c), Error, true)
	}

	return variable
}
