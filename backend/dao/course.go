package dao

import "time"

type Course struct {
	ID           int64     // Course ID
	Title        string    // Course title
	Description  string    // Course description
	Category     string    // Course Category. Allowed values: to be defined
	ImageURL     string    // Course image URL
	CreationDate time.Time // Course creation date
	LastUpdated  time.Time // Course last updated date
}
