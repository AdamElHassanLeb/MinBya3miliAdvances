package Middleware

import (
	"context"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var jwtKey = []byte(Env.GetString("JWT_KEY", "")) // Replace with your secret key

// ValidateJWT function checks the JWT token
func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &Services.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims and add to request context
		if claims, ok := token.Claims.(*Services.Claims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "phone_number", claims.PhoneNumber)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
