package clients

import (
	"backend/dao"
	"fmt"
	"strings"
	"time"
)

const (
	tableNameUsers         = "users"
	tableNameCourses       = "courses"
	tableNameSubscriptions = "subscriptions"
)

var (
	db = map[string][]interface{}{
		tableNameUsers: {
			dao.User{
				ID:           1,
				Email:        "emikohmann@gmail.com",
				PasswordHash: "5f4dcc3b5aa765d61d8327deb882cf99",
				Type:         "admin",
				CreationDate: time.Now().UTC(),
				LastUpdated:  time.Now().UTC(),
			},
			dao.User{
				ID:           2,
				Email:        "juanperez@gmail.com",
				PasswordHash: "5f4dcc3b5aa765d61d8327deb882cf99",
				Type:         "normal",
				CreationDate: time.Now().UTC(),
				LastUpdated:  time.Now().UTC(),
			},
		},
		tableNameCourses: {
			dao.Course{
				ID:           1,
				Title:        "Golang Programming Language",
				Description:  "Introduction course for the Golang programming language",
				Category:     "Programming",
				CreationDate: time.Now().UTC(),
				LastUpdated:  time.Now().UTC(),
			},
		},
		tableNameSubscriptions: {},
	}
)

func ConnectDatabase() error {
	// TODO: Connect to database here
	return nil
}

func SelectUserByEmail(email string) (dao.User, error) {
	for _, user := range db[tableNameUsers] {
		if userDAO := user.(dao.User); userDAO.Email == email {
			return userDAO, nil
		}
	}
	return dao.User{}, fmt.Errorf("not found user with email: %s", email)
}

func SelectCoursesWithFilter(query string) ([]dao.Course, error) {
	results := make([]dao.Course, 0)
	for _, course := range db[tableNameCourses] {
		if courseDAO := course.(dao.Course); query == "" ||
			strings.Contains(courseDAO.Title, query) ||
			strings.Contains(courseDAO.Description, query) {
			results = append(results, courseDAO)
		}
	}
	return results, nil
}

func SelectCourseByID(id int64) (dao.Course, error) {
	for _, course := range db[tableNameCourses] {
		if courseDAO := course.(dao.Course); courseDAO.ID == id {
			return courseDAO, nil
		}
	}
	return dao.Course{}, fmt.Errorf("not found course with ID: %d", id)
}

func InsertSubscription(userID int64, courseID int64) error {
	for _, subscription := range db[tableNameSubscriptions] {
		if subscriptionDAO := subscription.(dao.Subscription); subscriptionDAO.UserID == userID &&
			subscriptionDAO.CourseID == courseID {
			return fmt.Errorf("user %d is already subscribed to course %d", userID, courseID)
		}
	}

	db[tableNameSubscriptions] = append(db[tableNameSubscriptions], dao.Subscription{
		ID:           int64(len(db[tableNameSubscriptions]) + 1),
		UserID:       userID,
		CourseID:     courseID,
		CreationDate: time.Now().UTC(),
		LastUpdated:  time.Now().UTC(),
	})

	return nil
}
