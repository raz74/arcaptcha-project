package authentication

import (
	"Arc/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var _ = godotenv.Load(".env")
var secret = []byte(os.Getenv("SECRET"))

func GenerateToken(id int) (string, error) {
	expirationTime := time.Now().Add(24 * 7 * time.Hour)
	claims := &model.JwtCustomClimes{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(secret)
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
	
	if !token.Valid {
		return echo.ErrUnauthorized
	}

	claims, ok := token.Claims.(*model.JwtCustomClimes)
	if !ok {
		return echo.ErrUnauthorized
	}
	Id := claims.ID
	println(Id)

	return nil
}
