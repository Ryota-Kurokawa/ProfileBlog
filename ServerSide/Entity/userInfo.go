package entity

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model

	name       string   `gorm:"column:name"`
	age        int      `gorm:"column:age"`
	skillTag   []string `gorm:"column:skill_tag"`
	profileUrl []string `gorm:"column:profile_url"`
	profile    string   `gorm:"column:profile"`
}
