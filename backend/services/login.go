package services

import (
	"backend/clients"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
)

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

	// TODO: Replace this with JWT token generation
	token := hash

	return token, nil
}
