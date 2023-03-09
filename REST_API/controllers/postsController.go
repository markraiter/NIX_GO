package controllers

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"example.com/REST_API/initializers"
	"example.com/REST_API/models"
	_ "example.com/REST_API/cmd/docs"
	_ "github.com/swaggo/swag"
	_ "github.com/swaggo/echo-swagger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Getting posts godoc
// @Summary Get all posts
// @Description Get all posts from the database.
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Post
// @Failure 500 {string} string "Error fetching post"
// @Router /posts [get]
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


// Creating post godoc
// @Summary Create a new post
// @Description Create a new post with the specified data.
// @Accept  json
// @Produce  json
// @Param post body models.Post true "Post data"
// @Success 201 {object} models.Post
// @Failure 400 {string} string "Invalid post data"
// @Failure 500 {string} string "Error creating post"
// @Router /posts [post]
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


// Getting post by id godoc
// @Summary Get a post by ID
// @Description Get a post from the database by ID.
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {string} string "Post not found"
// @Router /posts/{id} [get]
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


// Updating post godoc
// @Summary Update a post by ID
// @Description Update a post in the database by ID.
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Param post body models.Post true "Post data"
// @Success 200 {object} models.Post
// @Failure 400 {string} string "Invalid post data"
// @Failure 404 {string} string "Post not found"
// @Failure 500 {string} string "Error updating post"
// @Router /posts/{id} [put]
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


// Deleting post godoc
// @Summary Delete a post by ID
// @Description Delete a post from the database by ID.
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 204 "Post deleted successfully"
// @Failure 404 {string} string "Post not found"
// @Failure 500 {string} string "Error deleting post"
// @Router /posts/{id} [delete]
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