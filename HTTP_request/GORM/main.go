package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	baseURL      = "https://jsonplaceholder.typicode.com"
	userID       = 7
	dbDriverName = "mysql"
)

type Post struct {
	gorm.Model
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	gorm.Model
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		"127.0.0.1:3306",
		"nix_beginner")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	err = db.AutoMigrate(&Post{}, &Comment{})
	if err != nil {
		log.Fatal(err)
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
			c, err := getComments(int(post.ID))
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
			err := db.Create(&comment).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Printf("Finished inserting comments for user %d\n", userID)
}

func getPosts(userID int, db *gorm.DB) ([]Post, error) {
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
		err := db.Create(&post).Error
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
