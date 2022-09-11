package main

//just running enough for complete migration
import (
	"apiexample/initializers"
	"apiexample/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
