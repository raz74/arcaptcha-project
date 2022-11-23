package main

import (
	"net/http"

	"Arc/model"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	dsn := "host=localhost user=admin password=123456 dbname=postgres port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db = database
	if err != nil {
		println(err)
		return
	}

	Db.AutoMigrate(&model.User{})

	e := echo.New()
	e.POST("/signup", Signup)

	e.Logger.Fatal(e.Start(":3000"))
}

func Signup(c echo.Context) error {

	var req userRequest
	c.Bind(&req)

	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	_ = Db.Create(&user)

	return c.JSON(http.StatusOK, user)
}
