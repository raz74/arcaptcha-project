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
	var req request.UserRequest
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

	return c.JSON(http.StatusOK, user.ToResponse())
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	var req request.UserRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	id := c.Param("id")
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return echo.ErrNotFound
	}
	err = u.repo.UpdateUser(id)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, user.ToResponse())
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	_ , err := u.repo.GetUserByID(id)
	if err != nil {
		return echo.ErrNotFound
	}

	err = u.repo.DeleteUser(id)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, id+" Deleted")
}
