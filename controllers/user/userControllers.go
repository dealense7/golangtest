package user

import (
	"github.com/dealense7/documentSignatures/requests"
	userRequest "github.com/dealense7/documentSignatures/requests/user"
	userService "github.com/dealense7/documentSignatures/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindItems(c *gin.Context) {
	items, err := userService.FindItems()

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	item, err := userService.FindByIdOrFail(id)

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Find Item to Update
	item, err := userService.FindByIdOrFail(id)

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
	err = userService.Update(item, request)

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Find Item to Update
	item, err := userService.FindByIdOrFail(id)

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err.GetMessage())
		return
	}

	// Delete
	userService.Delete(item)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Item Deleted",
	})
	return
}
