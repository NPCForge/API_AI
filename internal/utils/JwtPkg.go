package utils

import (
	"fmt"
	"my-api/config"
	"my-api/internal/models"
	"my-api/pkg"
	"time"

	"github.com/golang-jwt/jwt"
)

// CustomClaims defines custom JWT claims with a user ID.
type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT creates a JWT token containing the user ID and an expiration time of 7 days.
func GenerateJWT(userID string) (string, error) {
	var jwtSecret = []byte(config.GetEnvVariable("JWT_SECRET_KEY"))

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Expires in 7 days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// VerifyJWT verifies the provided JWT token and returns its claims if valid.
func VerifyJWT(tokenString string) (*models.CustomClaims, error) {
	var jwtSecret = []byte(config.GetEnvVariable("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	_, isStored := pkg.GetUserID(tokenString)
	if !isStored {
		return nil, fmt.Errorf("invalid token: not found in token store")
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GetUserIDFromJWT extracts and returns the user ID from a valid JWT token.
func GetUserIDFromJWT(tokenString string) (string, error) {
	var jwtSecret = []byte(config.GetEnvVariable("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok {
		return "", fmt.Errorf("token claims are not of type *models.CustomClaims")
	}

	return claims.UserID, nil
}
