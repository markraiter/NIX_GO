package controllers

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"example.com/REST_API/initializers"
	"example.com/REST_API/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Getting posts
func GetPosts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		responseType := c.QueryParam("type")
		if responseType != "xml" {
			responseType = "json"
		}

		posts := []models.Post{}
		if err := db.Preload("Comments").Find(&posts).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Error fetching posts")
		}

		if responseType == "xml" {
			xmlResponse, err := xml.MarshalIndent(posts, "", " ")
			if err != nil {
				return c.String(http.StatusInternalServerError, "Error encoding XML")
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			return c.Blob(http.StatusOK, "application/xml", xmlResponse)
		} else {
			c.Response().Header().Set("Content-Type", "application/json")
			return c.JSON(http.StatusOK, posts)
		}
	}
}


// Creating post
func CreatePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		post := new(models.Post)
		if err := c.Bind(post); err != nil {
			return c.String(http.StatusBadRequest, "Invalid post data")
		}

		if err := initializers.DB.Create(post).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Error creating post")
		}

		return c.JSON(http.StatusCreated, post)
	}
}


// Getting post by id
func GetPost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid post id")
		}

		post := new(models.Post)
		if err := initializers.DB.Preload("Comments").First(post, postID).Error; err != nil {
			return c.String(http.StatusNotFound, "Post not found")
		}

		return c.JSON(http.StatusOK, post)
	}
}


// Updating post
func UpdatePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid post ID")
		}

		post := new(models.Post)
		if err := initializers.DB.First(post, postID).Error; err != nil {
			return c.String(http.StatusNotFound, "Post not found")
		}

		if err := c.Bind(post); err != nil {
			return c.String(http.StatusBadRequest, "Invalid post data")
		}

		if err := initializers.DB.Save(post).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Error updating post")
		}

		return c.JSON(http.StatusOK, post)
	}
}


// Deleting post
func DeletePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid post ID")
		}

		post := new(models.Post)
		if err := initializers.DB.First(post, postID).Error; err != nil {
			return c.String(http.StatusNotFound, "Post not found")
		}

		if err := initializers.DB.Delete(post).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Error deleting post")
		}

		return c.NoContent(http.StatusNoContent)
	}
}