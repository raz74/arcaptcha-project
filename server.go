package main

import (
	"Arc/handlers"
	"Arc/model"
	"Arc/repository"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	repository.Initialize()
	e := echo.New()
	e.POST("/admmin/signup", handlers.Signup)
	e.POST("/admin/login", handlers.Login)

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading environment")
	}

	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClimes{},
		SigningKey: []byte(os.Getenv("SECRET")),
	}
	userGroup := e.Group("/users")
	websiteGroup := e.Group("/website")
	userGroup.Use(middleware.JWTWithConfig(config))
	
	userGroup.GET("/", handlers.GetAllUsers)
	userGroup.GET("/:id", handlers.GetUser)
	userGroup.POST("/", handlers.CreateUser)
	userGroup.PUT("/:id", handlers.UpdateUser)
	userGroup.DELETE("/:id", handlers.DeleteUser)
	
	websiteGroup.POST("/", handlers.CreateWebsite)
	websiteGroup.GET("/", handlers.GetAllWebsites)
	websiteGroup.GET("/:id", handlers.GetWebsite)
	websiteGroup.PUT("/:id", handlers.UpdateWebsite)
	websiteGroup.DELETE("/:id", handlers.DeleteWebsite)
	e.Logger.Fatal(e.Start(":3000"))
}
