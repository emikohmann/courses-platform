package database

import (
	"backend/dao"
	"backend/services/hashing"
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

	if err := db.AutoMigrate(&dao.User{}, &dao.Course{}, &dao.Subscription{}); err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return nil
}

func SelectUserByID(id int64) (dao.User, error) {
	var user dao.User
	if result := db.First(&user, id); result.Error != nil {
		return dao.User{}, fmt.Errorf("not found user with ID: %d", id)
	}
	return user, nil
}

func SelectUserByEmail(email string) (dao.User, error) {
	var user dao.User
	if result := db.Where("email = ?", email).First(&user); result.Error != nil {
		return dao.User{}, fmt.Errorf("not found user with email: %s", email)
	}
	return user, nil
}

func InsertUser(email string, password string, _type string) error {
	user := dao.User{
		Email:        email,
		PasswordHash: hashing.Hash(password),
		Type:         _type,
		CreationDate: time.Now().UTC(),
		LastUpdated:  time.Now().UTC(),
	}

	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func SelectCoursesWithFilter(query string) ([]dao.Course, error) {
	var courses []dao.Course
	if result := db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&courses); result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func SelectCourseByID(id int64) (dao.Course, error) {
	var course dao.Course
	if result := db.First(&course, id); result.Error != nil {
		return dao.Course{}, fmt.Errorf("not found course with ID: %d", id)
	}
	return course, nil
}

func SelectSubscriptionByUserAndCourse(userID int64, courseID int64) (dao.Subscription, error) {
	var subscription dao.Subscription
	if result := db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&subscription); result.Error != nil {
		return dao.Subscription{}, fmt.Errorf("not found subscription for user %d and course %d", userID, courseID)
	}
	return subscription, nil
}

func SelectSubscriptionsByUser(userID int64) ([]dao.Subscription, error) {
	var subscriptions []dao.Subscription
	if result := db.Where("user_id = ?", userID).Find(&subscriptions); result.Error != nil {
		return nil, result.Error
	}
	return subscriptions, nil
}

func InsertSubscription(userID int64, courseID int64) error {
	subscription := dao.Subscription{
		UserID:       userID,
		CourseID:     courseID,
		CreationDate: time.Now().UTC(),
		LastUpdated:  time.Now().UTC(),
	}

	if result := db.Create(&subscription); result.Error != nil {
		return result.Error
	}

	return nil
}
