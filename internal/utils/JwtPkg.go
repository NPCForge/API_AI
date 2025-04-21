package utils

import (
	"fmt"
	"my-api/config"
	"my-api/internal/models"
	"my-api/pkg"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID string) (string, error) {
	var jwtSecret = []byte(config.GetEnvVariable("JWT_SECRET_KEY"))

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add((time.Hour * 24) * 7).Unix(), // Expire dans 24h
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

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

	// Check if token is stored in connected tokens
	_, isStored := pkg.GetUserID(tokenString)

	if !isStored {
		return nil, fmt.Errorf("invalid token")
	}

	// Vérifie si le token est valide et les claims
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GetUserIDFromJWT(tokenString string) (string, error) {
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
		return "failed", err
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok {
		return "failed", fmt.Errorf("token claims are not of type *models.CustomClaims")
	}

	userID, err := strconv.Atoi(claims.UserID)

	if err != nil {
		return "failed", fmt.Errorf("failed to convert user_id to int: %v", err)
	}

	return strconv.Itoa(userID), nil
}
