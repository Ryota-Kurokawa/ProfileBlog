package main

import (
	"fmt"
	"net/http"
	"os"

	entity "github.com/Ryota-Kurokawa/ProfileBlog.git/Entity"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := dbInit()

	// Create
	db.AutoMigrate(&entity.UserInfo{})
	db.AutoMigrate(&entity.Blog{})

	user := entity.UserInfo{
		Name:    "Ryota",
		Age:     25,
		Profile: "I am a software engineer.",
	}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)

	blog := entity.Blog{
		Title:    "Title",
		Content:  "Content",
		AuthorId: uuid.New().String(),
	}
	result = db.Create(&blog)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)

	// ClientSideのindex.htmlを表示
	// http.Handle("/", http.FileServer(http.Dir("./ClientSide")))

	// ClientSideのindex.htmlを表示
	http.Handle("/", http.FileServer(http.Dir("./ClientSide")))
	http.ListenAndServe(":8080", nil)

	dbi, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer dbi.Close()
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
