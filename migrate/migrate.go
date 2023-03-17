package main

import (
	"backend/initializers"
	"backend/models"
)

func init() {
	initializers.ConnectToDb()
	initializers.LoadEnvvariable()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
