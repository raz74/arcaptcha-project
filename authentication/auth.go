package authentication

import (
	"Arc/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)


func GenerateToken(id int) (string, error) {
	expirationTime := time.Now().Add(24 * 7 * time.Hour)
	claims := &model.JwtCustomClimes{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateToken(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.ErrUnauthorized
	}
	claims, ok := token.Claims.(*model.JwtCustomClimes)
	if !ok {
		return echo.ErrUnauthorized
	}
	Id := claims.ID
	println(Id)

	// TODO: validate

	return nil
}
