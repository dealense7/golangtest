package main

import (
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/dealense7/documentSignatures/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Permission{},
		&models.Role{},
		&models.User{},
	)
}
