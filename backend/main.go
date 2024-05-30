package main

import (
	"backend/clients"
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	clients.ConnectDatabase()
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
	router.POST("/login", controllers.Login)
	router.GET("/courses/search", controllers.Search)
	router.GET("/courses/:id", controllers.Get)
	router.POST("/subscriptions", controllers.Subscribe)
	router.Run(":8080")
}
