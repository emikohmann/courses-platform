package users

import (
	"backend/domain/shared"
	users2 "backend/domain/users"
	"backend/services/users"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request users2.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	token, err := users.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, shared.Result{
			Message: fmt.Sprintf("Unauthorized login: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, users2.LoginResponse{
		Token: token,
	})
}
