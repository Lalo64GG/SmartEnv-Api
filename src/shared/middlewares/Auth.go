package middlewares

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("mi_clave_secreta") // Cambia esto por una clave secreta segura

// GenerateJWT genera un token JWT con los claims proporcionados
func GenerateJWT(userID int64, email string) (string, error) {
    // Definir los claims
    claims := CustomClaims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), 
            IssuedAt:  jwt.NewNumericDate(time.Now()),                     
            NotBefore: jwt.NewNumericDate(time.Now()),                     
            Issuer:    "myapp",                                         
        },
    }

    // Crear el token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Firmar el token con la clave secreta
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}