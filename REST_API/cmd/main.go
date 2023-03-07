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
	e.GET("/posts", controllers.GetPosts(initializers.DB))
	e.GET("/posts/:id", controllers.CreatePost(initializers.DB))
	e.POST("/posts", controllers.GetPosts(initializers.DB))
	e.PUT("/posts/:id", controllers.UpdatePost(initializers.DB))
	e.DELETE("/posts/:id", controllers.DeletePost(initializers.DB))

	//Operations with comments
	e.POST("/posts/:postId/comments", controllers.CreateComment(initializers.DB))
	e.PUT("/comments/:id", controllers.UpdateComment(initializers.DB))
	e.DELETE("/comments/:id", controllers.DeleteComment(initializers.DB))

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
