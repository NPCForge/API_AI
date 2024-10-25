package pkg

import (
	"my-api/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID string) (string, error) {
	var jwtSecret = []byte(config.GetEnvVariable("JWT_SECRET_KEY"))

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add((time.Hour * 24) * 7).Unix(), // Expire dans 24h
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
