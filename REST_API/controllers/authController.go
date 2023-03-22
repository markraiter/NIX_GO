package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"example.com/REST_API/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	_ "example.com/REST_API/cmd/docs"
	_ "github.com/swaggo/swag"
    _ "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func generateToken(u *models.User) (string, error) {
	claims := jwt.MapClaims {
		"id": u.ID,
		"username": u.Username,
		"email": u.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// @Summary Create a new user account
// @Tags Authentification
// @Description Register a new user account
// @Accept json
// @Produce json
// @Param user body object true "User object"
// @Param token body string true "JWT token"
// @Success 200 {string} string "JWT token"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal server error"
// @Router /registration [post]
func Registration(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string {
				"error": err.Error(),
			})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string {
				"error": err.Error(),
			})
		}

		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string {
				"error": err.Error(),
			})
		}

		token, err := generateToken(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string {
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string {
			"token": token,
		})
	}
}

// @Summary Login
// @Tags Authentification
// @Description logs in user with email and password
// @Accept  json
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {string} string "token"
// @Failure 400 {object} error "Bad request"
// @Failure 401 {object} error "Unauthorized"
// @Failure 500 {object} error "Internal server error"
// @Router /login [post]
func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string {
				"error": err.Error(),
			})
		}

		var dbUser models.User
		if err := db.Where("email = ?", user.Email).First(dbUser).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string {
				"error": "Invalid input!",
			})
		}

		token, err := generateToken(&dbUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string {
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string {
			"token": token,
		})
	}
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string {
				"error": "Authorization header missing",
			})
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte("secret"), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string {
				"error": err.Error(),
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["id"].(float64))
			c.Set("userID", userID)

			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string {
				"error": "Invalid token",
			})
		}
	}
}