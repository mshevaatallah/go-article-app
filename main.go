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
	r.GET("/article", middleware.RequireAuth, controllers.Profile)
	r.GET("/myprofile", middleware.RequireAuth, controllers.GetArticleByUser)
	r.GET("/article/:tag", middleware.RequireAuth, controllers.GetByTag)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)
	r.POST("/article", middleware.RequireAuth, controllers.CreateArticle)
	r.PUT("/article/:id", middleware.RequireAuth, controllers.UpdateArticle)

	r.Run()

}
