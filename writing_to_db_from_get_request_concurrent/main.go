package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// type Posts struct {
// 	UserId int `json:"userId"`
// 	Id int `json:"id"`
// 	Title string `json:"title"`
// 	Body string `json:"body"`
// }

// type Comments struct {
// 	PostId int `json:"postId"`
// 	Id int `json:"id"`
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// 	Body string `json:"body"`
// }

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
	
	resp1, err := http.Get(urlPosts)
	if err != nil {panic(err)}
	defer resp1.Body.Close()
	body1, _ := io.ReadAll(resp1.Body)
	dataPosts := string(body1)
	fmt.Println(dataPosts)
	// request dataComments from URL
	urlComments := "https://jsonplaceholder.typicode.com/comments?postId=7"
	
	resp2, err := http.Get(urlComments)
	if err != nil {panic(err)}
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	dataComments := string(body2)
	fmt.Println(dataComments)
}