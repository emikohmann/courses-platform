package services

import (
	"backend/clients"
	"backend/domain"
	"fmt"
	"strings"
)

func Search(query string) ([]domain.Course, error) {
	trimmed := strings.TrimSpace(query)

	courses, err := clients.SelectCoursesWithFilter(trimmed)
	if err != nil {
		return nil, fmt.Errorf("error getting courses from DB: %w", err)
	}

	results := make([]domain.Course, 0)
	for _, course := range courses {
		results = append(results, domain.Course{
			ID:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Category:     course.Category,
			CreationDate: course.CreationDate,
			LastUpdated:  course.LastUpdated,
		})
	}
	return results, nil
}

func Get(id int64) (domain.Course, error) {
	course, err := clients.SelectCourseByID(id)
	if err != nil {
		return domain.Course{}, fmt.Errorf("error getting course from DB: %w", err)
	}
	return domain.Course{
		ID:           course.ID,
		Title:        course.Title,
		Description:  course.Description,
		Category:     course.Category,
		CreationDate: course.CreationDate,
		LastUpdated:  course.LastUpdated,
	}, nil
}

func Subscribe(userID int64, courseID int64) error {
	if err := clients.InsertSubscription(userID, courseID); err != nil {
		return fmt.Errorf("error creating subscription in DB: %w", err)
	}
	return nil
}
