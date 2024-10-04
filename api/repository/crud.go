package repository

import (
	entity "github.com/Ryota-Kurokawa/ProfileBlog.git/entity"
	"gorm.io/gorm"
)

// 全てのユーザーを取得するためのインターフェース
type IUserRepository interface {
	GetAllUsers(users *[]entity.UserInfo) error
}

// UserRepository is a struct for UserRepository
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is a constructor for UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAllUsers implements IUserRepository.
func (r *UserRepository) GetAllUsers(users *[]entity.UserInfo) error {
	if err := r.db.Find(users).Error; err != nil {
		return err
	}
	return nil
}
