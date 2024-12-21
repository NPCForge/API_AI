package models

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
