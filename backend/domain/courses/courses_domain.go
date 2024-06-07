package courses

import "time"

type Course struct {
	ID           int64     `json:"id"`            // Course ID
	Title        string    `json:"title"`         // Course title
	Description  string    `json:"description"`   // Course description
	Category     string    `json:"category"`      // Course Category. Allowed values: to be defined
	ImageURL     string    `json:"image_url"`     // Course image URL
	CreationDate time.Time `json:"creation_date"` // Course creation date
	LastUpdated  time.Time `json:"last_updated"`  // Course last updated date
}

type SearchResponse struct {
	Results []Course `json:"results"`
}
