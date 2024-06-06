package courses

import (
	courses2 "backend/domain/courses"
	"backend/domain/shared"
	"backend/services/courses"
	"net/http"
	"strings"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := strings.TrimSpace(c.Query("query"))
	results, err := courses.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, shared.Result{
			Message: fmt.Sprintf("Error in search: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, courses2.SearchResponse{
		Results: results,
	})
}

func Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.Result{
			Message: fmt.Sprintf("Invalid ID: %s", err.Error()),
		})
		return
	}

	course, err := courses.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, shared.Result{
			Message: fmt.Sprintf("Error in get: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}

func Subscribe(c *gin.Context) {
	var request courses2.SubscribeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	if err := courses.Subscribe(request.UserID, request.CourseID); err != nil {
		c.JSON(http.StatusConflict, shared.Result{
			Message: fmt.Sprintf("Error in subscribe; %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, shared.Result{
		Message: fmt.Sprintf("User %d successfully subscribed to course %d", request.UserID, request.CourseID),
	})
}
