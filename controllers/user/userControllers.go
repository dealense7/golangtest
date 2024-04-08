package user

import (
	userServiceContract "github.com/dealense7/documentSignatures/app/contracts/services/user"
	"github.com/dealense7/documentSignatures/validation/requests"
	"github.com/dealense7/documentSignatures/validation/requests/v1/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UsersController struct {
	service userServiceContract.UserServiceContract
}

func NewUserController() *UsersController {
	service := userServiceContract.NewUserService()
	return &UsersController{service: service}
}

func (uc *UsersController) FindItems(c *gin.Context) {
	items, err := uc.service.FindItems()

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (uc *UsersController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	item, err := uc.service.FindByIdOrFail(id)

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func (uc *UsersController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Find Item to Update
	item, err := uc.service.FindByIdOrFail(id)

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	// Validate
	var request userRequest.UpdateUserRequest

	bindErr := c.ShouldBind(&request)
	if bindErr != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": requests.GenerateValidationErrors(request, bindErr),
		})
		return
	}

	// Update
	err = uc.service.Update(item, request)

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func (uc *UsersController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Find Item to Update
	item, err := uc.service.FindByIdOrFail(id)

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	// Delete
	uc.service.Delete(item)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Item Deleted",
	})
	return
}
