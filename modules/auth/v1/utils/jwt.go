package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type TokenDecoded struct {
	UserID primitive.ObjectID `json:"user_id"`
	Email  string             `json:"email"`
}

func GenerateToken(userID string, email string) (string, error) {
	// Create the claims
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key" // Default secret for development
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*TokenDecoded, error) {
	// Mocked validation
	userIdString := "68069217e5b753eed822d5d1" // Example ObjectID string
	userId, err := primitive.ObjectIDFromHex(userIdString)
	if err != nil {
		return nil, fmt.Errorf("error converting string to ObjectID: %v", err)
	}
	return &TokenDecoded{
		// convert string to ObjectID
		UserID: userId,
		Email:  "user@example.com",
	}, nil
}
