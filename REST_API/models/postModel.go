package models

import (
	"gorm.io/gorm"
)


type Post struct {
	gorm.Model
	UserID int    `json:"userId" xml:"userId"`
	Title  string `json:"title" xml:"title"`
	Body   string `json:"body" xml:"body"`
	Comments []Comment `json:"comments" xml:"comments"`
}

