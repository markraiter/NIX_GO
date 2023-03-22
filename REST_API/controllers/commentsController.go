package controllers

import (
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

// Creating comment godoc
// @Summary Create a new comment for a post
// @Description Create a new comment for a post with the specified data.
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param post_id path int true "Post ID"
// @Param comment body models.Comment true "Comment data"
// @Security ApiKeyAuth
// @Success 201 {object} models.Comment
// @Failure 400 {string} string "Invalid comment data"
// @Failure 404 {string} string "Post not found"
// @Failure 500 {string} string "Error creating comment"
// @Router /api/v1/posts/{post_id}/comments [post]
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


// Updating comment godoc
// @Summary Update a comment by ID
// @Description Update a comment in the database by ID.
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Param comment body models.Comment true "Comment data"
// @Security ApiKeyAuth
// @Success 200 {object} models.Comment
// @Failure 400 {string} string "Invalid comment data"
// @Failure 404 {string} string "Comment not found"
// @Failure 500 {string} string "Error updating comment"
// @Router /api/v1/comments/{id} [put]
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


// Deleting comment godoc
// @Summary Delete a comment by ID
// @Description Delete a comment from the database by ID.
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Security ApiKeyAuth
// @Success 204 "Comment deleted successfully"
// @Failure 404 {string} string "Comment not found"
// @Failure 500 {string} string "Error deleting comment"
// @Router /api/v1/comments/{id} [delete]
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