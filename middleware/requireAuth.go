package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-article-auth/initializers"
	"go-article-auth/models"
	"os"
	"time"
)

func IsAuth() gin.HandlerFunc {
	return RequireAuth(false)
}
func IsAdmin() gin.HandlerFunc {
	return RequireAuth(true)
}

func RequireAuth(check bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT string from the cookie
		tokenString, err := c.Cookie("Auth")
		if err != nil {
			c.JSON(401, gin.H{
				"status":  "error",
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(claims["exp"].(float64)) < float64(time.Now().Unix()) {
				c.JSON(401, gin.H{
					"status":  "error",
					"message": "token expired",
				})
				c.Abort()
				return
			}
			var user models.User
			initializers.DB.First(&user, claims["sub"])
			if user.ID == 0 {
				c.JSON(401, gin.H{
					"status":  "error",
					"message": "user not found",
				})
				c.Abort()
				return
			}
			c.Set("user", user)
			c.Set("user_id", user.ID)
			c.Set("role", user.Role)
			if check && user.Role != true {
				c.JSON(401, gin.H{
					"status":  "error",
					"message": "unauthorized request",
				})
				c.Abort()
				return
			}
		} else {
			fmt.Println(err)
		}
		c.Next()
	}
}
