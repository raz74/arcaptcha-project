package main

import (
	"Arc/handlers"
	"Arc/model"
	"Arc/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	repository.Initialize()
	e := echo.New()
	e.POST("/signup", handlers.Signup)
	e.POST("/login", handlers.Login)

	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClimes{},
		SigningKey: []byte("secret"),
	}

	userGroup := e.Group("/users")
	userGroup.Use(middleware.JWTWithConfig(config))

	userGroup.GET("/", handlers.GetAllUsers)
	userGroup.GET("/:id", handlers.GetUser)
	userGroup.PUT("/:id", handlers.UpdateUser)
	userGroup.DELETE("/:id", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(":3000"))
}
