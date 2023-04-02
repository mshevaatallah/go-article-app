package main

import (
	"github.com/gin-gonic/gin"
	"go-article-auth/controllers"
	"go-article-auth/initializers"
	"go-article-auth/middleware"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
	initializers.SyncDatabases()
}
func main() {
	r := gin.Default()
	r.POST("/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/article", middleware.IsAuth(), controllers.Profile)
	r.GET("/myprofile", middleware.IsAuth(), controllers.GetArticleByUser)
	r.GET("/article/:tag", middleware.IsAuth(), controllers.GetByTag)
	r.POST("/logout", middleware.IsAuth(), controllers.Logout)
	r.POST("/article", middleware.IsAuth(), controllers.CreateArticle)
	r.PUT("/article/:id", middleware.IsAuth(), controllers.UpdateArticle)
	r.DELETE("/article/:id", middleware.IsAdmin(), controllers.AdminDelete)

	r.Run()

}
