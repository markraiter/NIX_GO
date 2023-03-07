package controllers

import (
	"net/http"
	"strconv"

	"example.com/REST_API/initializers"
	"example.com/REST_API/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Creating comment
func CreateComment(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("postId"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid post ID")
		}

		post := new(models.Post)
		if err := initializers.DB.First(post, postID).Error; err != nil {
			return c.String(http.StatusNotFound, "Post not found")
		}

		comment := new(models.Comment)
		if err := c.Bind(comment); err != nil {
			return c.String(http.StatusBadRequest, "Invalid comment data")
		}

		comment.PostID = int(post.ID)
		if err := initializers.DB.Create(comment).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Error creating comment")
		}

		post.Comments = append(post.Comments, *comment)
		return c.JSON(http.StatusCreated, comment)
	}
}


// Updating comment
func UpdateComment(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		commentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid comment ID")
		}

		comment := new(models.Comment)
		if err := initializers.DB.First(comment, commentID).Error; err != nil {
			return c.String(http.StatusNotFound, "Comment not found")
		}

		if err := c.Bind(comment); err != nil {
			return c.String(http.StatusBadRequest, "Invalid comment data")
		}

		if err := initializers.DB.Save(comment).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Error updating comment")
		}

		return c.JSON(http.StatusOK, comment)
	}
}


// Deleting comment
func DeleteComment(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		commentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid comment ID")
		}

		comment := new(models.Comment)
		if err := initializers.DB.First(comment, commentID).Error; err != nil {
			return c.String(http.StatusNotFound, "Comment not found")
		}

		if err := initializers.DB.Delete(comment).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Error deleting comment")
		}

		return c.NoContent(http.StatusNoContent)
	}
}