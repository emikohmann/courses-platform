package courses

import (
	domain "backend/domain/courses"
	"backend/domain/shared"
	service "backend/services/courses"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.Result{
			Message: fmt.Sprintf("Invalid ID: %s", err.Error()),
		})
		return
	}

	course, err := service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, shared.Result{
			Message: fmt.Sprintf("Error in get: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}

func Search(c *gin.Context) {
	query := strings.TrimSpace(c.Query("query"))
	results, err := service.Search(query)
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
