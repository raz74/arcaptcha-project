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
	// result := db.Find(&users)
	err := repository.Db.Find(&users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorHandeling{
			Status: http.StatusBadRequest,
			Msg: "Bad request",
			Date: time.Now(),
			})}
	
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	var user model.User
	// User ID from path `users/:id`
	id := c.Param("id")
	err := repository.Db.Where("id = ?", id).Find(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorHandeling{
			Status: http.StatusBadRequest,
			Msg: "Bad request",
			Date: time.Now(),
		})}

	return c.JSON(http.StatusOK, user)
}

// select name, id from users
// select * from users where id = 3;


func DeleteUser(c echo.Context) error {
	var user model.User
	id := c.Param("id")
 	err := repository.Db.Where("id = ?", id).Find(&user).Error
	
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorHandeling{
			Status: http.StatusNotFound,
			Msg: "user not found",
			Date: time.Now(),
		})
	}

 	repository.Db.Delete(&user)

	return c.JSON(http.StatusOK, id+"Deleted",)
}
