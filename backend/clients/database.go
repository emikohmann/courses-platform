package clients

import (
	"backend/dao"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

func ConnectDatabase() error {
	dsn := "root:root@tcp(localhost:3306)/coursesplatform?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.AutoMigrate(&dao.User{}, &dao.Course{}, &dao.Subscription{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return nil
}

func SelectUserByID(id int64) (dao.User, error) {
	var user dao.User
	result := db.First(&user, id)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("not found user with ID: %d", id)
	}
	return user, nil
}

func SelectUserByEmail(email string) (dao.User, error) {
	var user dao.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("not found user with email: %s", email)
	}
	return user, nil
}

func SelectCoursesWithFilter(query string) ([]dao.Course, error) {
	var courses []dao.Course
	result := db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func SelectCourseByID(id int64) (dao.Course, error) {
	var course dao.Course
	result := db.First(&course, id)
	if result.Error != nil {
		return dao.Course{}, fmt.Errorf("not found course with ID: %d", id)
	}
	return course, nil
}

func InsertSubscription(userID int64, courseID int64) error {
	var subscription dao.Subscription
	result := db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&subscription)
	if result.Error == nil {
		return fmt.Errorf("user %d is already subscribed to course %d", userID, courseID)
	}

	subscription = dao.Subscription{
		UserID:       userID,
		CourseID:     courseID,
		CreationDate: time.Now().UTC(),
		LastUpdated:  time.Now().UTC(),
	}

	result = db.Create(&subscription)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
