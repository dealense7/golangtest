package main

import (
	models2 "github.com/dealense7/documentSignatures/app/models"
	"github.com/dealense7/documentSignatures/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models2.Permission{},
		&models2.Role{},
		&models2.User{},
	)
}
