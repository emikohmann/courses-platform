package domain

import "time"

type Course struct {
	ID           int64     `json:"id"`            // Course ID
	Title        string    `json:"title"`         // Course title
	Description  string    `json:"description"`   // Course description
	Category     string    `json:"category"`      // Course Category. Allowed values: to be defined
	CreationDate time.Time `json:"creation_date"` // Course creation date
	LastUpdated  time.Time `json:"last_updated"`  // Course last updated date
}

type SearchRequest struct {
	Query string `json:"query"`
}

type SearchResponse struct {
	Results []Course `json:"results"`
}

type SubscribeRequest struct {
	UserID   int64 `json:"user_id"`
	CourseID int64 `json:"course_id"`
}
