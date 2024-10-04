package main

import (
	"fmt"
	"os"

	entity "github.com/Ryota-Kurokawa/ProfileBlog.git/entity"
	repository "github.com/Ryota-Kurokawa/ProfileBlog.git/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	engine := gin.Default()

	db := dbInit()
	db.AutoMigrate(&entity.UserInfo{})
	db.AutoMigrate(&entity.Blog{})
	ur := repository.NewUserRepository(db)

	engine.GET("/users", func(c *gin.Context) {
		var users []entity.UserInfo
		if err := ur.GetAllUsers(&users); err != nil {
			c.JSON(500, gin.H{"status": "Internal Server Error"})
			return
		}
		c.JSON(200, users)
	})

	engine.Run(":8080")
}

func dbInit() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("データベースの接続に失敗しました")
	}
	return db
}
