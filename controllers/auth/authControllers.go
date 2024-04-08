package authController

import (
	userServiceContract "github.com/dealense7/documentSignatures/app/contracts/services/user"
	"github.com/dealense7/documentSignatures/app/v1/services/auth"
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/dealense7/documentSignatures/validation/requests"
	"github.com/dealense7/documentSignatures/validation/requests/v1/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	service userServiceContract.UserServiceContract
}

func NewAuthController() *AuthController {
	service := userServiceContract.NewUserService()
	return &AuthController{service: service}
}

func (ac *AuthController) SignUp(c *gin.Context) {

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
	item, createErr := ac.service.Create(request)

	if createErr != nil {
		c.AbortWithStatusJSON(createErr.GetCode(), createErr.GetMessage())
		return
	}

	// return
	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func (ac *AuthController) Auth(c *gin.Context) {
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
	tokenString, authErr := authService.Authorize(request)
	if authErr != nil {
		c.AbortWithStatusJSON(authErr.GetCode(), authErr.GetMessage())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func (ac *AuthController) GetMe(c *gin.Context) {
	user := initializers.AuthUser

	c.JSON(http.StatusOK, gin.H{
		"item": user,
	})
}
