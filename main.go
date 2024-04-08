package main

import (
	authController "github.com/dealense7/documentSignatures/controllers/auth"
	"github.com/dealense7/documentSignatures/controllers/user"
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/dealense7/documentSignatures/validation/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	route := r.Group("/api", middleware.SetApiVersion)

	// Initial Controllers
	authenticationController := authController.NewAuthController()
	UserController := user.NewUserController()

	// Define prefix
	userRoutes := route.Group("users")

	userRoutes.POST("/auth", authenticationController.Auth)
	userRoutes.POST("/register", authenticationController.SignUp)

	// Apply Middleware to route group
	userRoutesMiddleware := userRoutes.Group("/", middleware.RequireAuth)

	userRoutesMiddleware.GET("/me", authenticationController.GetMe)

	userRoutesMiddleware.GET("/", UserController.FindItems)
	userRoutesMiddleware.GET("/show/:id", UserController.Show)
	userRoutesMiddleware.PUT("/:id", UserController.Update)
	userRoutesMiddleware.DELETE("/:id", UserController.Delete)

	r.Run()
}
