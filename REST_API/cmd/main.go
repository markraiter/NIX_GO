package main

import (
	"os"

	"example.com/REST_API/controllers"
	"example.com/REST_API/initializers"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	e := echo.New()

	// Operations with posts
	e.POST("/posts", controllers.PostsCreate)
	e.PUT("/posts/:id", controllers.PostsUdate)
	e.GET("/posts", controllers.PostsIndex)
	e.GET("/posts/:id", controllers.PostsShow)
	e.DELETE("/posts/:id", controllers.PostsDelete)

	//Operations with comments
	

	e.Start(os.Getenv("SERVER_PORT"))
}
