package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/REST_API/initializers"
	"gorm.io/gorm"
)

const (
	BaseURL      = "https://jsonplaceholder.typicode.com"
	UserID       = 7
)

type Post struct {
	gorm.Model
	UserID int    `json:"userId" xml:"userId"`
	Title  string `json:"title" xml:"title"`
	Body   string `json:"body" xml:"body"`
}

func GetPosts(UserID int, db *gorm.DB) ([]Post, error)  {
	url := fmt.Sprintf("%s/posts?userId=%d", BaseURL, UserID)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var posts []Post
	err = json.NewDecoder(res.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		// Save post to database
		err := initializers.DB.Create(&post).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	return posts, nil
}
