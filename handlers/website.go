package handlers

import (
	"Arc/model"
	"net/http"
	"Arc/repository"
	"github.com/labstack/echo"
)

 func CreateWebsite (c echo.Context) error {
	var req model.CreateWebsiteRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	
	NewWebSite := &model.WebSite{
		User_Id: req.User_Id,
		Site_Key: req.Site_Key,
		Secret_Key: req.Secret_Key,
		Label: req.Label,
	}
	err:= repository.Db.Create(&NewWebSite).Error
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, NewWebSite)
}