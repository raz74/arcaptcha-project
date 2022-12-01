package handlers

import (
	"Arc/model"
	"Arc/repository"
	"Arc/handlers/request"
	"net/http"

	"github.com/labstack/echo"
)

func AddUserPlan( c echo.Context) error {
	var req request.PlanRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	var count int64
	repository.Db.Model(&model.UserPlan{}).Where("user_id= ?", req.UserId).Count(&count)
	if count >= 1 {
		return c.JSON(http.StatusBadRequest, "user already has a plan!")
	}

	NewUserPlan := &model.UserPlan{
		UserID: req.UserId,
		PlanID: req.PlanId,
		ExTime: req.ExTime,
	}

	err := repository.Db.Create(NewUserPlan).Error
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, NewUserPlan)
}

func GetUserPlan(c echo.Context) error {
	var userPlan model.UserPlan

	id := c.Param("user_id")
	result := repository.Db.Where("user_id = ?", id).First(&userPlan)
	if result.Error != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, userPlan)
}

func UpdateUserPlan(c echo.Context) error {
	var req request.PlanRequest

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	var userPlan model.UserPlan
	id := c.Param("user_id")
	result := repository.Db.Where("user_id = ?", id).First(&userPlan)
	if result.Error != nil {
		return echo.ErrNotFound
	}

	userPlan.PlanID = req.PlanId

	repository.Db.Save(&userPlan)
	return c.JSON(http.StatusOK, userPlan)
}

func DeleteUserPlan(c echo.Context) error {
	var userPlan model.UserPlan
	id := c.Param("user_id")
	result := repository.Db.Where("user_id", id).First(&userPlan)
	if result.Error != nil {
		return echo.ErrNotFound
	}

	repository.Db.Delete(&userPlan)
	return c.JSON(http.StatusOK, id+" userid's plan deleted")
}