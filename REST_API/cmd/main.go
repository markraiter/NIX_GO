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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()

	// Serve the Swagger documentation
	swagHandler := echoSwagger.WrapHandler
	e.GET("/swagger/*", swagHandler)

	// Group of routes that require authentification
	authGroup := e.Group("/api")
	authGroup.Use(controllers.JWTMiddleware)

	/////////////////ENDPOINTS/////////////////

	//Registration
	e.POST("/registration", controllers.Registration(initializers.DB))

	//Login
	e.POST("/login", controllers.Login(initializers.DB))

	// Operations with posts
	authGroup.GET("/v1/posts", controllers.GetPosts(initializers.DB))
	authGroup.GET("/v1/posts/:id", controllers.CreatePost(initializers.DB))
	authGroup.POST("/v1/posts", controllers.GetPosts(initializers.DB))
	authGroup.PUT("/v1/posts/:id", controllers.UpdatePost(initializers.DB))
	authGroup.DELETE("/v1/posts/:id", controllers.DeletePost(initializers.DB))

	//Operations with comments
	authGroup.POST("/v1/posts/:postId/comments", controllers.CreateComment(initializers.DB))
	authGroup.PUT("/v1/comments/:id", controllers.UpdateComment(initializers.DB))
	authGroup.DELETE("/v1/comments/:id", controllers.DeleteComment(initializers.DB))

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}