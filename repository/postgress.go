package repository

import (
	"Arc/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	// UpdateUser(user *model.User) error
	// GetUserByUsername(username string) (*model.User, error)
	// DeleteUser(id int) error
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) error {
	return u.Db.Create(user).Error
}

func (u *UserRepositoryImpl) GetUserByID(id string) (*model.User, error) {
	var user *model.User
	err := u.Db.Where("id = ?", id).First(user).Error
	return user, err
}

func (u *UserRepositoryImpl) UpdateUser(user *model.User) error {
	return nil
}