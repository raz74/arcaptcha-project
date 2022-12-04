package repository

import (
	"Arc/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	// UpdateUser(user *model.User) error
	// GetUserID(id int) (*model.User, error)
	// GetUserByUsername(username string) (*model.User, error)
	// DeleteUser(id int) error
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) error {
	return u.Db.Create(user).Error
}


