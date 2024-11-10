package auth

import (
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = Env.GetString("JWT_KEY", "") // Replace with an environment variable or secure key management

// GenerateJWT generates a new JWT token for a user.
func GenerateJWT(userID int) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   string(userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24-hour expiration
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
