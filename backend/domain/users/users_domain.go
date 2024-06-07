package users

import "time"

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Type         string    `json:"type"`
	CreationDate time.Time `json:"creation_date"`
	LastUpdated  time.Time `json:"last_updated"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
