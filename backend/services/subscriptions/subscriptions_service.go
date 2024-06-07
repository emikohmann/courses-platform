package subscriptions

import (
	"backend/clients/database"
	"backend/domain/subscriptions"
	"fmt"
)

func Create(userID int64, courseID int64) error {
	if _, err := database.SelectUserByID(userID); err != nil {
		return fmt.Errorf("error getting user from DB: %w", err)
	}

	if _, err := database.SelectCourseByID(courseID); err != nil {
		return fmt.Errorf("error getting course from DB: %w", err)
	}

	if _, err := database.SelectSubscriptionByUserAndCourse(userID, courseID); err == nil {
		return fmt.Errorf("subscription already exists for user %d and course %d", userID, courseID)
	}

	if err := database.InsertSubscription(userID, courseID); err != nil {
		return fmt.Errorf("error creating subscription in DB: %w", err)
	}

	return nil
}

func GetByUserID(userID int64) ([]subscriptions.Subscription, error) {
	records, err := database.SelectSubscriptionsByUser(userID)
	if err != nil {
		return nil, fmt.Errorf("error getting records from DB: %w", err)
	}

	results := make([]subscriptions.Subscription, 0)
	for _, record := range records {
		results = append(results, subscriptions.Subscription{
			ID:           record.ID,
			UserID:       record.UserID,
			CourseID:     record.CourseID,
			CreationDate: record.CreationDate,
			LastUpdated:  record.LastUpdated,
		})
	}
	return results, nil
}
