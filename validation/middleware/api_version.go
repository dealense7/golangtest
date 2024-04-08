package middleware

import (
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetApiVersion(c *gin.Context) {

	version, _ := strconv.Atoi(c.Request.Header.Get("X-Api-Version"))

	ApiSupportedVersions := []int{initializers.ApiV1}
	if !isInSlice(version, ApiSupportedVersions) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid api version",
		})
	}

	initializers.SetApiVersion(version)

	c.Next()
}

func isInSlice(value int, slice []int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
