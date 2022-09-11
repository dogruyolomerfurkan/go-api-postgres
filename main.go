package main

import (
	"apiexample/controllers"
	"apiexample/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.Run()
}
