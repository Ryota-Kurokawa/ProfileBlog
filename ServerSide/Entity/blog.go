package entity

import "gorm.io/gorm"

type Blog struct {
	gorm.Model

	Title    string
	Content  string
	AuthorId string
}
