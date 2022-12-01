package handlers

import (
	"Arc/authentication"
	"Arc/handlers/request"
	"Arc/model"
	"Arc/repository"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {

	var req request.SignupRequest

	if err := c.Bind(&req); err != nil {
		return err
	}
	hash, _ := hashPassword(req.Password)
		
	admin := model.Admin{
		Name:     req.Name,
		Password: hash,
		Email:    req.Email,
	}
	
	repository.CreateAdmin(&admin)
	return c.JSON(http.StatusOK, admin)
}

func Login(c echo.Context) error {
	var request request.LoginRequest

	if err := c.Bind(&request); err != nil {
		return err
	}
	
	var admin model.Admin

	result := repository.Db.Where("name = ?", request.Name).Find(&admin)
	if result.Error!= nil {
		return echo.ErrNotFound
	}

	// dbHash := repository.Db.Where("name = ?", request.Name).Find(admin.Password)
	match := checkPasswordHash(request.Password, admin.Password)
	if match != nil {
		return c.JSON(http.StatusUnauthorized, "Password is wrong! Try again.")
	}

	token, err := authentication.GenerateToken(admin.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(Password, hash string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
    return err 
}