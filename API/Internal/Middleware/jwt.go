package Middleware

import (
	"context"
	"errors"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

var jwtKey = []byte(Env.GetString("JWT_KEY", "")) // Replace with your secret key

// AuthMiddleware to verify the JWT and pass the user_id to the controller
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// The token should be in the form "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Parse the token
		claims := &Services.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Store user_id in the context to be used in the controller
		ctx := context.WithValue(r.Context(), "token_user_id", claims.UserID)

		// Pass the context with user_id to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
