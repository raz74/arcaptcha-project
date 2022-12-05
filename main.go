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
	e.POST("/admin/signup", handlers.Signup)
	e.POST("/admin/login", handlers.Login)

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading environment")
	}

	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClimes{},
		SigningKey: []byte(os.Getenv("SECRET")),
	}

	// e.Use(middleware.JWTWithConfig(config))
	addUserHandlers(e, config)
	addWebSiteHandlers(e, config)
	addPlanHanders(e, config)

	e.Logger.Fatal(e.Start(":3000"))
}

func addUserHandlers(e *echo.Echo, config middleware.JWTConfig) {
	repo := &repository.UserRepositoryImpl{
		Db: repository.Db,
	}

	h := handlers.NewUserHandler(repo)

	gp := e.Group("/users")
	gp.Use(middleware.JWTWithConfig(config))
	gp.GET("/", handlers.GetAllUsers)
	gp.GET("/:id", h.GetUser)
	gp.POST("/", h.CreateUser)
	gp.PUT("/:id", h.UpdateUser)
	gp.DELETE("/:id", h.DeleteUser)
}

func addWebSiteHandlers(e *echo.Echo, config middleware.JWTConfig) {
	gp := e.Group("/website")
	gp.Use(middleware.JWTWithConfig(config))
	gp.POST("/", handlers.CreateWebsite)
	gp.GET("/", handlers.GetAllWebsites)
	gp.GET("/:id", handlers.GetWebsite)
	gp.PUT("/:id", handlers.UpdateWebsite)
	gp.DELETE("/:id", handlers.DeleteWebsite)
}

func addPlanHanders(e *echo.Echo, config middleware.JWTConfig) {
	gp := e.Group("/user/plan")
	gp.Use(middleware.JWTWithConfig(config))
	gp.POST("/", handlers.AddUserPlan)
	gp.GET("/:user_id", handlers.GetUserPlan)
	gp.PUT("/:user_id", handlers.UpdateUserPlan)
	gp.DELETE("/:user_id", handlers.DeleteUserPlan)
}
