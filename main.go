package main

import (
	"backend/controllers"
	"backend/initializers"

	"github.com/gin-gonic/gin"
)

// this func run before main
func init() {
	initializers.LoadEnvvariable()
	initializers.ConnectToDb()
}

func main() {

	router := gin.Default()
	router.POST("/connect", controllers.PostCreat)
	router.Run()
}
