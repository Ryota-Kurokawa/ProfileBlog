package entity

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model

	Name    string
	Age     int
	Profile string
}
