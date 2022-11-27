package handelers

import (
	"Arc/model"
	"Arc/repository"
	"net/http"
	"time"
	"github.com/labstack/echo"
)

func GetAllUsers(c echo.Context) error {
	var users []model.User
	err := repository.Db.Find(&users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewHTTPEror(http.StatusBadRequest, "Bad request"))
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	var user model.User
	// User ID from path `users/:id`
	id := c.Param("id")
	err := repository.Db.Where("id = ?", id).Find(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.HTTPError{
			Status: http.StatusBadRequest,
			Msg:    "Bad request",
			Date:   time.Now(),
		})
	}

	return c.JSON(http.StatusOK, user)
}


func UpdateUser(c echo.Context) error {
	var req model.UserRequest
	if err:= c.Bind(&req) ; err != nil {
		return c.JSON(http.StatusBadRequest, model.HTTPError{
			Status: http.StatusBadRequest,
			Msg:    "Bad request",
			Date:   time.Now(),
		})
	}

	name := c.Param("name")
	email := c.Param("email")

	user := model.User{
        Name: req.Name,
		Email: req.Email,
	}

	err := repository.Db.Where("name=?", name).Find(&user)
	user.Email = email

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.HTTPError{
			Status: http.StatusBadRequest,
			Msg:    "Bad request",
			Date:   time.Now(),
		})
	}
	repository.Db.Save(&user)
	return c.JSON(http.StatusOK, name+"user successfully updated")
}


func DeleteUser(c echo.Context) error {
	var user model.User
	id := c.Param("id")
	err := repository.Db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return c.JSON(http.StatusNotFound, model.HTTPError{
			Status: http.StatusNotFound,
			Msg:    "user not found",
			Date:   time.Now(),
		})
	}

	repository.Db.Delete(&user)

	return c.JSON(http.StatusOK, id+"Deleted")
}


func CreateUser(c echo.Context) error {
	var req model.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.HTTPError{
			Status: http.StatusBadRequest,
			Msg:    "Invalid type of request",
			Date:   time.Now(),
		})
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
	
