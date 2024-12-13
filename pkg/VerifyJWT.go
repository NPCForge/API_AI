package pkg

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"my-api/config"
	"my-api/internal/models"
)

func VerifyJWT(tokenString string) (*models.CustomClaims, error) {
	var jwtSecret = []byte(config.GetEnvVariable("JWT_SECRET_KEY"))

	// Parse et vérifie le token
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Vérifie que la méthode de signature est bien HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Vérifie si le token est valide et les claims
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
