package controllers

import (
	"net/http"

	"example.com/REST_API/initializers"
	"example.com/REST_API/models"
	"github.com/labstack/echo/v4"
)

func PostsCreate(c echo.Context) error {
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body: body.Body,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		return c.String(http.StatusBadRequest, "")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"post": post})
}

func PostsIndex(c echo.Context) error {
	var posts []models.Post
	initializers.DB.First(&posts)

	return c.JSON(http.StatusOK, map[string]interface{}{"posts": posts})
}

func PostsShow(c echo.Context) error {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	return c.JSON(http.StatusOK, map[string]interface{} {"post": post})
}

func PostsUdate(c echo.Context) error {
	id := c.Param("id")
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	return c.JSON(http.StatusOK, map[string]interface{} {"post": post})
}

func PostsDelete(c echo.Context) error {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)

	return c.String(http.StatusOK, "")
}