package main

import (
	"backend/clients/database"
	"backend/controllers/cors"
	"backend/controllers/courses"
	"backend/controllers/subscriptions"
	"backend/controllers/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	if err := database.ConnectDatabase(); err != nil {
		fmt.Println(fmt.Sprintf("Error connecting database: %v", err))
		return
	}

	router := gin.New()
	router.Use(cors.AllowCORS)
	router.Use(func(c *gin.Context) {
		time.Sleep(200 * time.Millisecond)
		c.Next()
	})
	router.GET("/users/:id", users.Get)
	router.POST("/users/login", users.Login)
	router.POST("/users/signup", users.Signup)
	router.GET("/courses/search", courses.Search)
	router.GET("/courses/:id", courses.Get)
	router.POST("/subscriptions", subscriptions.Create)
	router.GET("/users/:id/subscriptions", subscriptions.GetByUserID)

	if err := router.Run(":8080"); err != nil {
		fmt.Println(fmt.Sprintf("Error running app: %v", err))
		return
	}
}
