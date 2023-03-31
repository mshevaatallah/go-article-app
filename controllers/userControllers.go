package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-article-auth/initializers"
	"go-article-auth/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	// take data from inputs body

	var body struct {
		Email    string
		Password string
		Username string
		Name     string
		Age      json.Number
	}
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "internal server error",
		})
		return
	}
	// save the user to database
	user := models.User{Email: body.Email, Password: string(hash), Username: body.Username, Name: body.Name, Age: body.Age}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "email user exists",
		})
		return
	}
	// respond
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user created",
	})

}

func Login(c *gin.Context) {
	//get data from user input or from body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}
	// find user by email and validate it

	var user models.User
	//query to database
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "user not found",
		})
		return
	}

	// compare password input with password in database(hashes)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "invalid password",
		})
		return

	}
	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "token validate error",
		})
		return
	}
	// respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(200, gin.H{
		"status": "success",
	})

}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func Profile(c *gin.Context) {
	// ambil data user dari context
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
