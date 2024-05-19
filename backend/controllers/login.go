package controllers

import (
	"backend/domain"
	"backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	token, err := services.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Result{
			Message: fmt.Sprintf("Unauthorized login: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		Token: token,
	})
}
