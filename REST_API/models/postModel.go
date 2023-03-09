package models

type Post struct {
	UserID   int       `json:"userId" xml:"userId"`
	ID       int       `json:"id" xml:"id"`
	Title    string    `json:"title" xml:"title"`
	Body     string    `json:"body" xml:"body"`
	Comments []Comment `json:"comments" xml:"comments"`
}

