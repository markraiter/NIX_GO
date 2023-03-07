package main

import (
	"os"

	"example.com/REST_API/controllers"
	"example.com/REST_API/initializers"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	_ "example.com/REST_API/cmd/docs"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// @Title Echo Swagger Example REST_API
// @Description Example REST_API for demonstrating Swagger with Echo framework.
// @Version 1.0.0.
// @Host localhost:8080
// @BasePath /
func main() {
	e := echo.New()

	// Serve the Swagger documentation
	swagHandler := echoSwagger.WrapHandler
	e.GET("/swagger/*", swagHandler)

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
