package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/REST_API/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetPosts(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/posts", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Create a mock database connection
    mockDB, mock, err := sqlmock.New()
    if err != nil {
        t.Fatal(err)
    }
    defer mockDB.Close()

    // Add an expectation for the SELECT VERSION() query
    mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.0"))

    // Create a GORM database connection using the mock database
    db, err := gorm.Open(mysql.New(mysql.Config{
        DriverName: "mysql",
        Conn:       mockDB,
    }), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }

    // Set up the mock query response
    rows := sqlmock.NewRows([]string{"id", "title", "content"}).AddRow(1, "Test Post", "This is a test post")
    mock.ExpectQuery("SELECT \\* FROM `posts`").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM `comments` WHERE `comments`.`post_id` = \\?").WillReturnRows(sqlmock.NewRows([]string{"id", "comment"}))

    // Call the GetPosts handler function
    err = GetPosts(db)(c)
    if err != nil {
        t.Fatal(err)
    }

    // Check the response status code
    if rec.Code != http.StatusOK {
        t.Errorf("expected status OK; got %v", rec.Code)
    }

    // Check the response content type
    if !strings.Contains(rec.Header().Get("Content-Type"), "application/json") {
        t.Errorf("expected response content type to be application/json; got %v", rec.Header().Get("Content-Type"))
    }

    // Check the response body
    var posts []models.Post
    err = json.Unmarshal(rec.Body.Bytes(), &posts)
    if err != nil {
        t.Fatal(err)
    }
    if len(posts) != 1 || posts[0].Title != "Test Post" {
        t.Errorf("unexpected response body; got %v", posts)
    }

    // Check that all expectations were met
    err = mock.ExpectationsWereMet()
    if err != nil {
        t.Error(err)
    }
}