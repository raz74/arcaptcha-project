package postgres

import (
	"Arc/model"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	UpdateUser(id string) error
	DeleteUser(id string) error
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) error {
	return u.Db.Create(user).Error
}

func (u *UserRepositoryImpl) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := u.Db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepositoryImpl) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	if err := u.Db.Find(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepositoryImpl) UpdateUser(id string) error {
	var user model.User
	err := u.Db.Save(&user).Error
	return err
}

func (u *UserRepositoryImpl) DeleteUser(id string) error {
	var user model.User
	var userPlan model.UserPlan
	if err := u.Db.Where("user_id = ?", id).Delete(&userPlan).Error; err != nil {
		return echo.ErrNotFound
	}
	if err := u.Db.Where("id= ?", id).Delete(&user).Error; err != nil {
		return echo.ErrNotFound
	}
	return nil
}
