package handelers

import (
	"Arc/model"
	"Arc/repository"
	"Arc/handelers/request"
	"net/http"
	"github.com/labstack/echo"
)

func GetAllUsers(c echo.Context) error {
	var users []model.User
	err := repository.Db.Find(&users)

	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	var user model.User
	// User ID from path `users/:id`
	id := c.Param("id")
	err := repository.Db.Where("id = ?", id).Find(&user)

	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, user)
}


func UpdateUser(c echo.Context) error {
	var req handelers.UserRequest
	if err:= c.Bind(&req) ; err != nil {
		return echo.ErrBadRequest
	}

	id := c.Param("id")

	user := model.User{}

	err := repository.Db.Where("id=?", id).Find(&user)
	user.Email = req.Email

	if err != nil {
		return echo.ErrBadRequest
	}
	repository.Db.Save(&user)
	return c.JSON(http.StatusOK, id+"user successfully updated")
}


func DeleteUser(c echo.Context) error {
	var user model.User
	id := c.Param("id")
	err := repository.Db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return echo.ErrNotFound
	}

	repository.Db.Delete(&user)

	return c.JSON(http.StatusOK, id+"Deleted")
}


func CreateUser(c echo.Context) error {
	var req handelers.UserRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	NewUser:= &model.User{
		Password: req.Password,
		Name: req.Name,
		ID : req.ID,
		Email: req.Email,
		Phone: req.Phone,
		Company_name: req.Company_name,
		Job_title: req.Job_title,
		Active: req.Active,	
	}
	return c.JSON(http.StatusOK, NewUser)
}	
	
