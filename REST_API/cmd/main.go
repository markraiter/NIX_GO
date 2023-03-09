package main

import (
	"os"

	"example.com/REST_API/controllers"
	"example.com/REST_API/initializers"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/swag"

	_ "example.com/REST_API/cmd/docs"
	_ "example.com/REST_API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// @title Echo Swagger Example REST_API
// @description Example REST_API for demonstrating Swagger with Echo framework.
// @version 1.0.0.
// @host localhost:8080
// @basePath /api/v1
func main() {
	e := echo.New()

	// Serve the Swagger documentation
	swagHandler := echoSwagger.WrapHandler
	e.GET("/swagger/*", swagHandler)

	// Operations with posts
	e.GET("/api/v1/posts", controllers.GetPosts(initializers.DB))
	e.GET("/api/v1/posts/:id", controllers.CreatePost(initializers.DB))
	e.POST("/api/v1/posts", controllers.GetPosts(initializers.DB))
	e.PUT("/api/v1/posts/:id", controllers.UpdatePost(initializers.DB))
	e.DELETE("/api/v1/posts/:id", controllers.DeletePost(initializers.DB))

	//Operations with comments
	e.POST("/api/v1/posts/:postId/comments", controllers.CreateComment(initializers.DB))
	e.PUT("/api/v1/comments/:id", controllers.UpdateComment(initializers.DB))
	e.DELETE("/api/v1/comments/:id", controllers.DeleteComment(initializers.DB))

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
