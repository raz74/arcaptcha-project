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

type AdminHandler struct {
	repo repository.AdminRepositiry
}

func NewAdminHandler(r repository.AdminRepositiry) *AdminHandler {
	return &AdminHandler{
		repo: r,
	}
}

func (u *AdminHandler) CreateAdmin(admin *model.Admin) error {
	// return repository.Db.Create(admin).Error
	err := u.repo.CreateAdmin(admin)
	if err != nil {
		return echo.ErrBadRequest
	}
	return nil
}

func (u *AdminHandler) Signup(c echo.Context) error {

	var req request.SignupRequest

	if err := c.Bind(&req); err != nil {
		return err
	}
	hash, _ := hashPassword(req.Password)
		
	Admin := model.Admin{
		Name:     req.Name,
		Password: hash,
		Email:    req.Email,
	}
	err := u.repo.ChechAdminEmailUnique(req.Email)
	if err != nil {
		return c.JSON(http.StatusForbidden, "This email is already exists. Try another!")
	}
	
	err = u.repo.CreateAdmin(&Admin)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, "New admin signup successfully.")
}

func (u *AdminHandler) Login(c echo.Context) error {
	var request request.LoginRequest

	if err := c.Bind(&request); err != nil {
		return err
	}
	
	admin, err := u.repo.Login(request.Email)
	if err != nil {
		return echo.ErrNotFound
	}

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