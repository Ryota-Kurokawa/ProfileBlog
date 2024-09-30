package main

import (
	"fmt"
	"os"

	entity "github.com/Ryota-Kurokawa/ProfileBlog.git/Entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := dbInit()

	// Create
	db.AutoMigrate(&entity.UserInfo{})
	db.AutoMigrate(&entity.Blog{})
}

func dbInit() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
