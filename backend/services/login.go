package services

import (
	"backend/clients"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var jwtSecret = []byte("your_secret_key") // Replace with your actual secret key

func generateJWT(email string) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token expiry time (72 hours)
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(jwtSecret)
}

func Login(email string, password string) (string, error) {
	if strings.TrimSpace(email) == "" {
		return "", errors.New("email is required")
	}

	if strings.TrimSpace(password) == "" {
		return "", errors.New("password is required")
	}

	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))

	userDAO, err := clients.SelectUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("error getting user from DB: %w", err)
	}

	if hash != userDAO.PasswordHash {
		return "", fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := generateJWT(email)
	if err != nil {
		return "", fmt.Errorf("error generating JWT token: %w", err)
	}

	return token, nil
}
