package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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

func getRequest(url string, /**data chan string**/) {
	resp, err := http.Get(url)
	if err != nil {panic(err)}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	// data <- string(body)
	// fmt.Println(string(body))
	writeToDB(string(body))
}

func writeToDB(values string) {
	//Connecting to DB
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/comments")
	if err != nil {panic(err)}
	defer db.Close()

	fmt.Println("Connetcted to MySQL")

	//Inserting data
	insert, err := db.Query("INSERT INTO `posts` (`userId`, `id`, `title`, `body`)", values)
	if err != nil {panic(err)}
	defer insert.Close()
}

func main() {
	urlPosts := "https://jsonplaceholder.typicode.com/posts?userId=7"
	// urlComments := "https://jsonplaceholder.typicode.com/comments?postId=7"
	// dataPosts := make(chan string)
	// dataComments := make(chan string)

	/**go**/ getRequest(urlPosts, /**dataPosts**/)
	// fmt.Println(<- dataPosts)
	// close(dataPosts)
	// go getRequest(urlComments, /**dataComments**/)
	// fmt.Println(<- dataComments)
	// close(dataComments)
}