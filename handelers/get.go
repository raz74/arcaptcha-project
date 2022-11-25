package handelers

import (
	"Arc/model"
	"net/http"
	"github.com/labstack/echo"
	"Arc/repository"
)

func GetAllUsers(c echo.Context) error {
	var users []model.User
	// result := db.Find(&users)
	_ = repository.Db.Find(&users)
	return c.JSON(http.StatusOK, users)
}


func GetUser(c echo.Context) error {
	var user model.User
	// User ID from path `users/:id`
 	id := c.Param("id")
	repository.Db.Where("id = ?", id).Find(&user)
 	return c.JSON(http.StatusOK, user)
}

// select name, id from users
// select * from users where id = 3;