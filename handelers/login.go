package handelers

import (
	"fmt"

	"Arc/model"

	"github.com/labstack/echo"

	"gorm.io/gorm"
)

var Db *gorm.DB

func Login(c echo.Context) error {
	var request model.LoginRequest
	println(request.Username)
	fmt.Printf("30 %+v\n", request)

	c.Bind(&request)

	println(request.Username)
	fmt.Printf("%+v\n", request)

	var user model.User
	fmt.Printf("%+v\n", user)
	Db.Where("name = ?", request.Username).First(&user)
	fmt.Printf("%+v\n", user)


	// if username != "admin" || password != "123456" {
	// 	return echo.ErrUnauthorized
	// }

	return nil
}