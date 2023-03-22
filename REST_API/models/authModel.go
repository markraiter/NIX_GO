package models

type User struct {
	ID       int    `json:"id" xml:"id"`
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}
