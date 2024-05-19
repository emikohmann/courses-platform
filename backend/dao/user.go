package dao

import "time"

type User struct {
	ID           int64     // User ID
	Email        string    // User Email
	PasswordHash string    // User Password Hash
	Type         string    // User Type. Allowed values: admin, normal
	CreationDate time.Time // User creation date
	LastUpdated  time.Time // User last updated date
}
