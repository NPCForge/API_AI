package config

import (
	"fmt"
	"my-api/internal/models"

	"github.com/go-resty/resty/v2"
)

var (
	clientPocketbase *resty.Client
	authToken        string
)

func InitClient() {
	clientPocketbase = resty.New()

	authUrl := "http://localhost:8090/api/admins/auth-with-password"

	// Requête pour obtenir le token d'authentification
	resp, err := clientPocketbase.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"identity": GetEnvVariable("EMAIL_POCKETBATE"),
			"password": GetEnvVariable("PASSWORD_POCKETBASE"),
		}).
		SetResult(&models.AuthResponse{}).
		Post(authUrl)

	if err != nil {
		fmt.Println("Erreur lors de la connexion :", err)
		return
	}

	// Extraire et stocker le token d'authentification
	authResp := resp.Result().(*models.AuthResponse)
	authToken = authResp.Token
	fmt.Println("Token reçu:", authToken)

	// Configure le client pour ajouter automatiquement le token à chaque requête
	clientPocketbase.SetAuthToken(authToken)
}

// GetClient retourne le client PocketBase configuré
func GetClient() *resty.Client {
	return clientPocketbase
}
