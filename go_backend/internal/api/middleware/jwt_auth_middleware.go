package middleware

import (
	"chess_log/go_backend/internal/helpers"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const ContextUserClaims = contextKey("userClaims")

func JWTAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the 'Authorization' header from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Authorization header required.")
			return
		}

		// Split the header into parts to separate the bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Invalid Authorization header format.")
			return
		}

		// Retrieve the JWT secret key from environment variables
		tokenStr := parts[1]
		secret := os.Getenv("JWT_SECRET_KEY")

		if secret == "" {
			helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Missing JWT secret key.")
			return
		}

		// Use MapClaims to parse
		claims := jwt.MapClaims{}

		// Parse the token with claims
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			// This callback function returns the secret key for token verification
			return []byte(secret), nil
		})
		// Check for parsing errors or if the token is not valid
		if err != nil || !token.Valid {
			helpers.RespondErrorJSON(w, http.StatusUnauthorized, err, "Invalid or expired token.")
			return
		}

		// Token is valid → save claims in context for handlers
		ctx := context.WithValue(r.Context(), ContextUserClaims, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
