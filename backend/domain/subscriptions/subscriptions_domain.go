package subscriptions

import "time"

type Subscription struct {
	ID           int64     `json:"id"`            // Subscription ID
	UserID       int64     `json:"user_id"`       // Course title
	CourseID     int64     `json:"course_id"`     // Course description
	CreationDate time.Time `json:"creation_date"` // Course creation date
	LastUpdated  time.Time `json:"last_updated"`  // Course last updated date
}

type CreateRequest struct {
	UserID   int64 `json:"user_id"`
	CourseID int64 `json:"course_id"`
}

type SearchResponse struct {
	Results []Subscription `json:"results"`
}
