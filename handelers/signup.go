package handelers

import (

	"Arc/model"
	"Arc/repository"

	"github.com/labstack/echo"

	"net/http"
)

func Signup(c echo.Context) error {

	var req model.UserRequest
	c.Bind(&req)

	user := model.User{
		Name: req.Username,
		Password: req.Password,
		// Email:    req.Email,
	}
	repository.CreateUser(&user)

	return c.JSON(http.StatusOK, user)
}