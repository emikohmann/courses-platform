package subscriptions

import (
	"backend/domain/shared"
	domain "backend/domain/subscriptions"
	service "backend/services/subscriptions"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Create(c *gin.Context) {
	var request domain.CreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	if err := service.Create(request.UserID, request.CourseID); err != nil {
		c.JSON(http.StatusConflict, shared.Result{
			Message: fmt.Sprintf("Error in subscribe; %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, shared.Result{
		Message: fmt.Sprintf("User %d successfully subscribed to course %d", request.UserID, request.CourseID),
	})
}

func GetByUserID(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.Result{
			Message: fmt.Sprintf("Invalid user ID: %s", err.Error()),
		})
		return
	}
	results, err := service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, shared.Result{
			Message: fmt.Sprintf("Error in search: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SearchResponse{
		Results: results,
	})
}
