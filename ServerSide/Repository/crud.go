package repository

import (
	"fmt"
	"log"

	entity "github.com/Ryota-Kurokawa/ProfileBlog.git/Entity"
	"gorm.io/gorm"
)

func Insert(db *gorm.DB) {
	user := entity.UserInfo{
		Name:    "Ryota",
		Age:     25,
		Profile: "I am a software engineer.",
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}
