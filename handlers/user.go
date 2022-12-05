package handlers

import (
	"Arc/handlers/request"
	"Arc/handlers/response"
	"Arc/model"
	"Arc/repository"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(r repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: r,
	}
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var req request.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	hash, _ := hashPassword(req.Password)

	NewUser := &model.User{
		Password:    hash,
		Name:        req.Name,
		Email:       req.Email,
		Phone:       req.Phone,
		CompanyName: req.CompanyName,
		JobTitle:    req.JobTitle,
		Active:      req.Active,
	}
	// err := repository.Db.Create(&NewUser).Error
	err := u.repo.CreateUser(NewUser)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, NewUser.ToResponse())
}

func GetAllUsers(c echo.Context) error {
	var users []response.UserResponse
	err := repository.Db.Find(&users).Error
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, users)
}

func (u *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return echo.ErrNotFound
	}

	response := user.ToResponse()

	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) error {
	var req response.UserResponse
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
	var user response.UserResponse
	id := c.Param("id")
	result := repository.Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return echo.ErrNotFound
	}

	err := repository.Db.Delete(&user).Error
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, id+"Deleted")
}
