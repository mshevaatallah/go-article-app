package initializers

import "go-article-auth/models"

func SyncDatabases() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Article{})
	DB.Model(&models.User{}).Association("Articles")
}
