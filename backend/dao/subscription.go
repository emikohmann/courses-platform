package dao

import "time"

type Subscription struct {
	ID           int64     // Subscription ID
	UserID       int64     // Subscription User ID
	CourseID     int64     // Subscription Course ID
	CreationDate time.Time // Subscription creation date
	LastUpdated  time.Time // Subscription last updated date
}
