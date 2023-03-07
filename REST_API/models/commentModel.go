package models

import (

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID int    `json:"postId" xml:"postId"`
	Name   string `json:"name" xml:"name"`
	Email  string `json:"email" xml:"email"`
	Body   string `json:"body" xml:"body"`
}