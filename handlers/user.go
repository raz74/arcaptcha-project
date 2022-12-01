package handlers

import (
	"Arc/authentication"
	"Arc/handlers/request"
	"Arc/model"
	"Arc/repository"
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	var req request.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	hash, _ := hashPassword(req.Password)

	NewUser := &model.User{
		Password:     hash,
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		Company_name: req.Company_name,
		Job_title:    req.Job_title,
		Active:       req.Active,
	}
	err := repository.Db.Create(&NewUser).Error
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, NewUser)
}

func GetAllUsers(c echo.Context) error {
	var users []model.User
	err := repository.Db.Find(&users).Error

	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	if err := authentication.ValidateToken(c); err != nil {
		return err
	}

	var user model.User

	id := c.Param("id")
	err := repository.Db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	var req request.SignupRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	id := c.Param("id")

	user := model.User{}

	err := repository.Db.Where("id=?", id).Find(&user).Error
	if err != nil {
		return echo.ErrBadRequest
	}
	user.Email = req.Email
	repository.Db.Save(&user)
	return c.JSON(http.StatusOK, id+"user successfully updated")
}

func DeleteUser(c echo.Context) error {
	var user model.User
	id := c.Param("id")
	result := repository.Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return echo.ErrNotFound
	}

	repository.Db.Delete(&user)

	return c.JSON(http.StatusOK, id+"Deleted")
}
