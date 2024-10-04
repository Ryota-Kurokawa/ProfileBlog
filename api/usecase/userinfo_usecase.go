package usecase

import (
	repository "github.com/Ryota-Kurokawa/ProfileBlog.git/Repository"
	"github.com/Ryota-Kurokawa/ProfileBlog.git/entity"
)

// IUserInfoUsecase is an interface for UserInfoUsecase
type IUserInfoUsecase interface {
	GetAllUsers(users *[]entity.UserInfo) error
}

// UserInfoUsecase is a struct for UserInfoUsecase
type UserInfoUsecase struct {
	ur repository.IUserRepository
}

// GetAllUsers implements IUserInfoUsecase.

// NewUserInfoUsecase is a constructor for UserInfoUsecase
// func NewUserInfoUsecase(ur repository.IUserRepository) IUserInfoUsecase {
// 	return &UserInfoUsecase{ur}
// }
