package courses

import (
	"backend/clients/database"
	"backend/domain/courses"
	"fmt"
	"strings"
)

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

func Subscribe(userID int64, courseID int64) error {
	if _, err := database.SelectUserByID(userID); err != nil {
		return fmt.Errorf("error getting user from DB: %w", err)
	}

	if _, err := database.SelectCourseByID(courseID); err != nil {
		return fmt.Errorf("error getting course from DB: %w", err)
	}

	if err := database.InsertSubscription(userID, courseID); err != nil {
		return fmt.Errorf("error creating subscription in DB: %w", err)
	}

	return nil
}
