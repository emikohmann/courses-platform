package users

import (
	"backend/clients/database"
	"backend/domain/users"
	"backend/services/hashing"
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

func Get(id int64) (users.User, error) {
	record, err := database.SelectUserByID(id)
	if err != nil {
		return users.User{}, fmt.Errorf("error getting record from DB: %w", err)
	}

	return users.User{
		ID:           record.ID,
		Email:        record.Email,
		Type:         record.Type,
		CreationDate: record.CreationDate,
		LastUpdated:  record.LastUpdated,
	}, nil
}

func Login(email string, password string) (string, error) {
	if strings.TrimSpace(email) == "" {
		return "", errors.New("email is required")
	}

	if strings.TrimSpace(password) == "" {
		return "", errors.New("password is required")
	}

	hash := hashing.Hash(password)

	userDAO, err := database.SelectUserByEmail(email)
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

func Signup(email string, password string, _type string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(password) == "" {
		return errors.New("password is required")
	}

	if strings.TrimSpace(_type) == "" {
		return errors.New("type is required")
	}

	if _type != "admin" && _type != "normal" {
		return errors.New("type must be admin or normal")
	}

	if _, err := database.SelectUserByEmail(email); err == nil {
		return fmt.Errorf("user already exists for email %s", email)
	}

	if err := database.InsertUser(email, password, _type); err != nil {
		return fmt.Errorf("error inserting user into DB: %w", err)
	}

	return nil
}
