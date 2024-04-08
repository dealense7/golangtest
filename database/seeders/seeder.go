package main

import (
	"github.com/dealense7/documentSignatures/database/seeders/seeds"
	"github.com/dealense7/documentSignatures/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	seeds.UserSeed()
	seeds.AclSeed()
}
