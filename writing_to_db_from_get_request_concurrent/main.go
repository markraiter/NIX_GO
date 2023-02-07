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

// Posts struct...
type Posts struct {
	UserId int
	Id int
	Title string
	Body string
}

// Comments struct...
type Comments struct {
	PostId int
	Id int
	Name string
	Email string
	Body string
}

func main() {
	// Install environment variables
	if err := godotenv.Load(); err != nil {fmt.Println(err.Error())}
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
	if pingErr != nil {fmt.Println(err.Error())}
	fmt.Println("Connection installed")

	// request dataPosts from URL
	urlPosts := "https://jsonplaceholder.typicode.com/posts?userId=7"
	urlComments := "https://jsonplaceholder.typicode.com/comments?postId=7"

	posts := getPosts(urlPosts)
	comments := getComments(urlComments)

	// Insert posts into DB

	insertPosts, err := db.Exec("INSERT INTO posts (userId, id, title, body) VALUES (?, ?, ?, ?)", posts.UserId, posts.Id, posts.Title, posts.Body)
	if err != nil {fmt.Println(err.Error())}
	rowsPosts, err := insertPosts.LastInsertId()
	if err != nil {fmt.Println(err.Error())}
	fmt.Println(rowsPosts)

	// Insert comments into DB
	insertComments, err := db.Exec("INSERT INTO comments (postId, id, name, email, body) VALUES (?, ?, ?, ?, ?)", comments.PostId, comments.Id, comments.Name, comments.Email, comments.Body)
	if err != nil {fmt.Println(err.Error())}
	rowsComments, err := insertComments.LastInsertId()
	if err != nil {fmt.Println(err.Error())}
	fmt.Println(rowsComments)
	
}

// Fetch posts into struct
func getPosts(url string) Posts {
	resp, err := http.Get(url)
	if err != nil {fmt.Println(err.Error())}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var p Posts
	parseError := json.Unmarshal(body, &p)
	if parseError != nil {fmt.Println(parseError.Error())}
	return p
}

// Fetch comments into struct
func getComments(url string) Comments {
	resp, err := http.Get(url)
	if err != nil {fmt.Println(err.Error())}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var c Comments
	parseError := json.Unmarshal(body, &c)
	if parseError != nil {fmt.Println(parseError.Error())}
	return c
}