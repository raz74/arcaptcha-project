package handelers

import (
	"Arc/model"
	"Arc/repository"
	"fmt"
	"net/http"
	"time"
	"Arc/handelers/request"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {

	var req handelers.UserRequest
	c.Bind(&req)

	if err := c.Bind(req); err != nil {
		fmt.Println(err.Error())
		panic("failed to signup")
	}

	user := model.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	repository.CreateUser(&user)

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	var request handelers.LoginRequest

	c.Bind(&request)

	var user model.User

	repository.Db.Where("name = ?", request.Name).First(&user)

	climes := &model.JwtCustomClimes{
		ID: request.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, climes)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
