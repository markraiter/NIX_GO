package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// var db *sql.DB

type Posts struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

type Comments struct {
	PostId int `json:"postId"`
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string `json:"body"`
}

func main() {
	// Install environment variables
	if err := godotenv.Load(); err != nil {panic(err)}
	//Capture connection properties
	cfg := mysql.Config {
		User: os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "nix_beginner",
	}
	
	// Get a database handle
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {panic(err)}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {panic(err)}
	fmt.Println("Connection installed")

	// request dataPosts from URL
	urlPosts := "https://jsonplaceholder.typicode.com/posts?userId=7"
	respPosts, err := http.Get(urlPosts)
	if err != nil {panic(err)}
	defer respPosts.Body.Close()
	bodyPosts, _ := io.ReadAll(respPosts.Body)

	// request dataComments from URL
	urlComments := "https://jsonplaceholder.typicode.com/comments?postId=7"
	respComments, err := http.Get(urlComments)
	if err != nil {panic(err)}
	defer respComments.Body.Close()
	bodyComments, _ := io.ReadAll(respComments.Body)

	// parse JSON to struct Posts
	dataPosts := Posts{}
	json.Unmarshal([]byte(bodyPosts), &dataPosts)
	if err != nil {panic(err)}

	// parse JSON to struct Comments
	dataComments := Comments{}
	json.Unmarshal([]byte(bodyComments), &dataComments)
	if err != nil {panic(err)}

	// write data (posts) to DB
	insertPosts, err := db.Query("INSERT INTO posts (userId, id, title, body) VALUES (?, ?, ?, ?)", dataPosts.UserId, dataPosts.Id, dataPosts.Title, dataPosts.Body)
	if err != nil {panic(err)}
	defer insertPosts.Close()

	// write data (comments) to DB
	insertComments, err := db.Query("INSERT INTO comments (postId, id, name, email, body) VALUES (?, ?, ?, ?, ?)", dataComments.PostId, dataComments.Id, dataComments.Name, dataComments.Email, dataComments.Body)
	if err != nil {panic(err)}
	defer insertComments.Close()
}