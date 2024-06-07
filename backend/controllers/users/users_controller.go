package users

import (
	"backend/domain/shared"
	domain "backend/domain/users"
	service "backend/services/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func Login(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	token, err := service.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, shared.Result{
			Message: fmt.Sprintf("Unauthorized login: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		Token: token,
	})
}

func Signup(c *gin.Context) {
	var request domain.SignupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	if err := service.Signup(request.Email, request.Password, request.Type); err != nil {
		c.JSON(http.StatusConflict, shared.Result{
			Message: fmt.Sprintf("Error in signup: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, shared.Result{
		Message: "User successfully created",
	})
}
