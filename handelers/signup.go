package handelers

import (
	"Arc/model"
	"Arc/repository"
	"fmt"

	"github.com/labstack/echo"

	"net/http"
)

func Signup(c echo.Context) error {

	var req model.UserRequest
	c.Bind(&req)

	if err := c.Bind(req); err != nil {
		fmt.Println(err.Error())
		panic("failed to signup")
	}

	user := model.User{
		Name: req.Name,
		Password: req.Password,
		// Email:    req.Email,
	}
	repository.CreateUser(&user)

	return c.JSON(http.StatusOK, user)
}