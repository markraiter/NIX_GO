package main

import (
	"example.com/REST_API/initializers"
	"example.com/REST_API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.Comment{})

}