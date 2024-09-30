package entity

import "gorm.io/gorm"

type Blog struct {
	gorm.Model

	title    string `gorm:"column:title"`
	content  string `gorm:"column:content"`
	authorId string `gorm:"column:author_id"`
}
