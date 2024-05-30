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
				Title:        "Python for Data Science, AI & Development",
				Description:  "What you'll learn Learn Python - the most popular programming language and for Data Science and Software Development.  Apply Python programming logic Variables, Data Structures, Branching, Loops, Functions, Objects & Classes.  Demonstrate proficiency in using Python libraries such as Pandas & Numpy, and developing code using Jupyter Notebooks.",
				Category:     "Programming",
				ImageURL:     "https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://s3.amazonaws.com/coursera-course-photos/fc/c1b8dfbac740999b6256aca490de43/Python-Image.jpg?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&fit=crop&q=50",
				CreationDate: time.Now().UTC(),
				LastUpdated:  time.Now().UTC(),
			},
			dao.Course{
				ID:           2,
				Title:        "Computer Science: Programming with a Purpose",
				Description:  "The basis for education in the last millennium was “reading, writing, and arithmetic;” now it is reading, writing, and computing. Learning to program is an essential part of the education of every student, not just in the sciences and engineering, but in the arts, social sciences, and humanities, as well.",
				Category:     "Programming",
				ImageURL:     "https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://s3.amazonaws.com/coursera-course-photos/86/9297002bbc11e8b82d4d350d2ce323/IntroCSlogo.png?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&fit=crop&q=50",
				CreationDate: time.Now().UTC(),
				LastUpdated:  time.Now().UTC(),
			},
			dao.Course{
				ID:           3,
				Title:        "Java Programming and Software Engineering Fundamentals Specialization",
				Description:  "Advance your subject-matter expertise. Learn in-demand skills from university and industry experts. Master a subject or tool with hands-on projects. Develop a deep understanding of key concepts. Earn a career certificate from Duke University.",
				Category:     "Programming",
				ImageURL:     "https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://coursera_assets.s3.amazonaws.com/collections/product-images/ibm-full-stack-cloud-developer.jpeg?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop",
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

func SelectUserByID(id int64) (dao.User, error) {
	for _, user := range db[tableNameUsers] {
		if userDAO := user.(dao.User); userDAO.ID == id {
			return userDAO, nil
		}
	}
	return dao.User{}, fmt.Errorf("not found user with ID: %d", id)
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
