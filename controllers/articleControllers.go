package controllers

import (
	"github.com/gin-gonic/gin"
	"go-article-auth/initializers"
	"go-article-auth/models"
)

func CreateArticle(c *gin.Context) {
	var body struct {
		Title string
		Desc  string
		Tag   string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}
	//userid := c.GetUint("user_id")
	//
	//fmt.Println(userid)

	article := models.Article{Title: body.Title, Tag: body.Tag, Desc: body.Desc, UserID: c.GetUint("user_id")}
	result := initializers.DB.Create(&article)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "save article fail",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user created",
	})
}

func GetByTag(c *gin.Context) {
	tag := c.Param("tag")
	items := []models.Article{}
	initializers.DB.Where("tag LIKE ?", "%"+tag+"%").Find(&items)
	c.JSON(200, gin.H{
		"status": "succes",
		"data":   items,
	})
}
