package courses

import (
	"backend/clients/database"
	"backend/domain/courses"
	"fmt"
	"strings"
)

func Get(id int64) (courses.Course, error) {
	record, err := database.SelectCourseByID(id)
	if err != nil {
		return courses.Course{}, fmt.Errorf("error getting record from DB: %w", err)
	}

	return courses.Course{
		ID:           record.ID,
		Title:        record.Title,
		Description:  record.Description,
		Category:     record.Category,
		ImageURL:     record.ImageURL,
		CreationDate: record.CreationDate,
		LastUpdated:  record.LastUpdated,
	}, nil
}

func Search(query string) ([]courses.Course, error) {
	trimmed := strings.TrimSpace(query)

	records, err := database.SelectCoursesWithFilter(trimmed)
	if err != nil {
		return nil, fmt.Errorf("error getting records from DB: %w", err)
	}

	results := make([]courses.Course, 0)
	for _, record := range records {
		results = append(results, courses.Course{
			ID:           record.ID,
			Title:        record.Title,
			Description:  record.Description,
			Category:     record.Category,
			ImageURL:     record.ImageURL,
			CreationDate: record.CreationDate,
			LastUpdated:  record.LastUpdated,
		})
	}
	return results, nil
}
