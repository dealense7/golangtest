package main

import (
	authController "github.com/dealense7/documentSignatures/controllers/auth"
	"github.com/dealense7/documentSignatures/controllers/user"
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/dealense7/documentSignatures/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	userRoutes := r.Group("users")

	userRoutes.POST("/auth", authController.Auth)
	userRoutes.POST("/register", authController.SignUp)

	// Apply Middleware to route group
	userRoutesMiddleware := userRoutes.Group("/", middleware.RequireAuth)

	userRoutesMiddleware.GET("/me", authController.GetMe)

	userRoutesMiddleware.GET("/", user.FindItems)
	userRoutesMiddleware.GET("/show/:id", user.Show)
	userRoutesMiddleware.PUT("/:id", user.Update)
	userRoutesMiddleware.DELETE("/:id", user.Delete)

	r.Run()
}
