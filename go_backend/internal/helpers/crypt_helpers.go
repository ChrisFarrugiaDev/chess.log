package helpers

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Generate a hashed password with a cost of bcrypt.DefaultCost (usually 10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	// Compare the hashed password with the plain-text password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // If no error, passwords match
}

func GenerateJWT(payload map[string]interface{}, exp time.Time) (string, error) {

	claims := jwt.MapClaims{}

	// // Copy payload fields into claims
	// for k, v := range payload {
	// 	claims[k] = v
	// }

	// Copy payload into JWT claims using modern Go function
	maps.Copy(claims, payload)

	// Add expiration
	claims["exp"] = exp.Unix()

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("JWT_EMAIL_SECRET"))

	// Sign token with secret
	signed, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signed, nil
}

func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	secret := []byte(os.Getenv("JWT_EMAIL_SECRET"))

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {

		// Ensure signing method is HMAC (HS256)
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		} else {
			return secret, nil
		}
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
