package main

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/joho/godotenv"

	"github.com/go-sql-driver/mysql"
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
}