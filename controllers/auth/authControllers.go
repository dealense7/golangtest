package authController

import (
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/dealense7/documentSignatures/requests"
	authRequest "github.com/dealense7/documentSignatures/requests/auth"
	userService "github.com/dealense7/documentSignatures/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {

	var request authRequest.UserRegistrationRequest

	// Validate
	err := c.ShouldBind(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": requests.GenerateValidationErrors(request, err),
		})
		return
	}

	// Create
	item, createErr := userService.Create(request)

	if createErr != nil {
		c.AbortWithStatusJSON(createErr.GetCode(), createErr.GetMessage())
		return
	}

	// return
	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func Auth(c *gin.Context) {
	var request authRequest.AuthRequest

	// Validate
	err := c.ShouldBind(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": requests.GenerateValidationErrors(request, err),
		})
		return
	}

	// Authorize
	tokenString, authErr := userService.Authorize(request)
	if authErr != nil {
		c.AbortWithStatusJSON(authErr.GetCode(), authErr.GetMessage())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func GetMe(c *gin.Context) {
	user := initializers.AuthUser

	c.JSON(http.StatusOK, gin.H{
		"item": user,
	})
}
