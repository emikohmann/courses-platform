package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	var request domain.SearchRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	results, err := services.Search(request.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Result{
			Message: fmt.Sprintf("Error in search: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SearchResponse{
		Results: results,
	})
}

func Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: fmt.Sprintf("Invalid ID: %s", err.Error()),
		})
		return
	}

	course, err := services.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Result{
			Message: fmt.Sprintf("Error in get: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}

func Subscribe(c *gin.Context) {
	var request domain.SubscribeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	if err := services.Subscribe(request.UserID, request.CourseID); err != nil {
		c.JSON(http.StatusConflict, domain.Result{
			Message: fmt.Sprintf("Error in subscribe; %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Result{
		Message: fmt.Sprintf("User %d successfully subscribed to course %d", request.UserID, request.CourseID),
	})
}
