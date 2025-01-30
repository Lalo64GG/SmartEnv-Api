package middlewares

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}