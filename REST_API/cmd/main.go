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
	postGroup := authGroup.Group("/v1/posts")
	commentsGroup := authGroup.Group("/v1/posts/:postId/comments")

	/////////////////ENDPOINTS/////////////////

	//Registration
	e.POST("/registration", controllers.Registration(initializers.DB))

	//Login
	e.POST("/login", controllers.Login(initializers.DB))

	// Operations with posts
	postGroup.GET("", controllers.GetPosts(initializers.DB))
	postGroup.GET("/:id", controllers.GetPost(initializers.DB))
	postGroup.POST("", controllers.CreatePost(initializers.DB))
	postGroup.PUT("/:id", controllers.UpdatePost(initializers.DB))
	postGroup.DELETE("/:id", controllers.DeletePost(initializers.DB))

	//Operations with comments
	commentsGroup.POST("", controllers.CreateComment(initializers.DB))
	commentsGroup.PUT("/:id", controllers.UpdateComment(initializers.DB))
	commentsGroup.DELETE("/:id", controllers.DeleteComment(initializers.DB))

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}