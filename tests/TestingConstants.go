package apiTesting

import (
	"fmt"
	"github.com/joho/godotenv"
	. "my-api/pkg"
	"os"
)

func getEnvVariable(c string) string {
	if err := godotenv.Load("../.env"); err != nil {
		DisplayContext("Error while loading .env", Error, true)
	}

	variable := os.Getenv(c)
	if variable == "" {
		DisplayContext(fmt.Sprintf("%s undefined in .env", c), Error, true)
	}

	return variable
}

const (
	HttpBaseUrl  = "http://localhost:3000/"
	WsConnectURL = "ws://localhost:3000/ws"
	TokenFile    = "token.json"

	Password = "Password"
	WsID     = "User_01_test_ws"
	HttpID   = "User_01_test_http"
)

var ApiKey = getEnvVariable("API_KEY_REGISTER")
