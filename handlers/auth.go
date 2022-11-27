package handlers

import (
	"Arc/authentication"
	"Arc/handlers/request"
	"Arc/model"
	"Arc/repository"
	"net/http"

	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {

	var req request.SignupRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	admin := model.Admin{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	repository.CreateAdmin(&admin)

	return c.JSON(http.StatusOK, admin)
}

func Login(c echo.Context) error {
	var request request.LoginRequest

	if err := c.Bind(&request); err != nil {
		return err
	}

	var admin model.Admin

	err := repository.Db.Where("name = ?", request.Name).First(&admin).Error
	if err != nil {
		return err
	}

	token, err := authentication.GenerateToken(admin.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
