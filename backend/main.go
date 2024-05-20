package main

import (
	"backend/clients"
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	clients.ConnectDatabase()
	router := gin.New()
	router.POST("/login", controllers.Login)
	router.POST("/courses/search", controllers.Search)
	router.GET("/courses/:id", controllers.Get)
	router.POST("/subscriptions", controllers.Subscribe)
	router.Run(":8080")
}
