package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID int    `json:"postId" xml:"postId"`
	Name   string `json:"name" xml:"name"`
	Email  string `json:"email" xml:"email"`
	Body   string `json:"body" xml:"body"`
}

func GetComments(postID int) ([]Comment, error) {
	url := fmt.Sprintf("%s/comments?postId=%d", BaseURL, postID)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var comments []Comment
	err = json.NewDecoder(res.Body).Decode(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}