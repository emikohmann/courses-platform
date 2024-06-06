package main

import (
	"backend/clients/database"
	"backend/controllers/courses"
	"backend/controllers/users"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	router := gin.New()
	// Middleware to handle CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	})
	router.POST("/login", users.Login)
	router.GET("/courses/search", courses.Search)
	router.GET("/courses/:id", courses.Get)
	router.POST("/subscriptions", courses.Subscribe)
	router.Run(":8080")
}
