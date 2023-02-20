package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	baseURL      = "https://jsonplaceholder.typicode.com"
	userID       = 7
	dbDriverName = "mysql"
)

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	// Install environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}
	//Capture connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "nix_beginner",
	}

	db, err := sql.Open(dbDriverName, cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = createTables(db)
	if err != nil {
		log.Println(err.Error())
	}

	posts, err := getPosts(userID, db)
	if err != nil {
		log.Println(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(len(posts))
	comments := make(chan []Comment, len(posts))

	for _, post := range posts {
		go func(post Post) {
			defer wg.Done()
			c, err := getComments(post.ID)
			if err != nil {
				log.Println(err.Error())
				return
			}
			comments <- c
		}(post)
	}

	go func() {
		wg.Wait()
		close(comments)
	}()

	for c := range comments {
		for _, comment := range c {
			_, err := db.Exec("INSERT INTO comments (post_id, id, name, email, body) VALUES (?, ?, ?, ?, ?)",
				comment.PostID, comment.ID, comment.Name, comment.Email, comment.Body)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Printf("Finished inserting comments for user %d\n", userID)
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			user_id INT NOT NULL,
			id INT NOT NULL PRIMARY KEY,
			title VARCHAR(255),
			body TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			post_id INT NOT NULL,
			id INT NOT NULL PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255),
			body TEXT,
			FOREIGN KEY (post_id) REFERENCES posts(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func getPosts(userID int, db *sql.DB) ([]Post, error) {
	url := fmt.Sprintf("%s/posts?userId=%d", baseURL, userID)
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
		_, err := db.Exec("INSERT INTO posts (id, user_id, title, body) VALUES (?, ?, ?, ?)",
			post.ID, post.UserID, post.Title, post.Body)
		if err != nil {
			fmt.Println(err)
		}
	}

	return posts, nil
}

func getComments(postID int) ([]Comment, error) {
	url := fmt.Sprintf("%s/comments?postId=%d", baseURL, postID)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var comments []Comment
	err = json.NewDecoder(res.Body).Decode(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil 
}
