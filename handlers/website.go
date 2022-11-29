package handlers

import (
	"Arc/model"
	"Arc/repository"
	"net/http"

	"github.com/labstack/echo"
)

func CreateWebsite(c echo.Context) error {
	var req model.CreateWebsiteRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	NewWebSite := &model.Website{
		UserId:    req.UserId,
		SiteKey:   req.SiteKey,
		SecretKey: req.SecretKey,
		Label:     req.Label,
		WebsiteV1: model.WebsiteV1{
			Level: req.Level,
		},
	}
	err := repository.Db.Create(&NewWebSite).Error
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, NewWebSite)
}
