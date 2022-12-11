package postgres

import (
	"Arc/model"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(admin *model.Admin) error
	CheckAdminEmailUnique(Email string) error
	GetByEmail(Email string) (*model.Admin, error)
}

type AdminRepositoryImpl struct {
	Db *gorm.DB
}

func (u *AdminRepositoryImpl) CreateAdmin(admin *model.Admin) error {
	return u.Db.Create(admin).Error
}

func (u *AdminRepositoryImpl) CheckAdminEmailUnique(Email string) error {
	var admin model.Admin
	err := u.Db.Where("email=?", Email).Find(&admin).RowsAffected
	if err > 0 {
		return echo.ErrForbidden
	}
	return echo.ErrBadRequest
}

func (u *AdminRepositoryImpl) GetByEmail(Email string) (*model.Admin, error) {
	var admin model.Admin
	err := u.Db.Where("email = ?", Email).Find(&admin).Error

	return &admin, err
}
