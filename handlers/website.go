package handlers

import (
	"Arc/handlers/request"
	"Arc/model"
	"Arc/repository"
	"net/http"

	"github.com/labstack/echo"
)

func CreateWebsite(c echo.Context) error {
	var req request.CreateWebsiteRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	NewWebSite := &model.Website{
		UserId:    req.UserId,
		SiteKey:   req.SiteKey,
		SecretKey: req.SecretKey,
		Label:     req.Label,	
	}
	
	err := repository.Db.Create(&NewWebSite).Error
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, NewWebSite)
}

func GetAllWebsites(c echo.Context) error {
	var websites []model.Website
	err := repository.Db.Preload("Domains").Preload("WebsiteV1").Find(&websites).Error
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, websites)
}

func GetWebsite(c echo.Context) error {
	var website model.Website

	id := c.Param("id")
	result := repository.Db.Where("id = ?", id).First(&website)
	if result.Error != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, website)
}

func UpdateWebsite(c echo.Context) error {
	var req request.UpdateWebsiteRequest

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	var website model.Website

	id := c.Param("id")
	result := repository.Db.Where("id = ?", id).First(&website)
	if result.Error != nil {
		return echo.ErrNotFound
	}

	website.SecretKey = req.SecretKey
	website.Label = req.Label
	website.SiteKey = req.SiteKey

	repository.Db.Save(&website)
	return c.JSON(http.StatusOK, website)
}

func DeleteWebsite(c echo.Context) error {
	var website model.Website
	id := c.Param("id")
	result := repository.Db.Where("id = ?", id).First(&website)
	if result.Error != nil {
		return echo.ErrNotFound
	}
	repository.Db.Delete(&website)
	return c.JSON(http.StatusOK, id+" deleted")
}
