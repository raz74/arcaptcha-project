package main

import (
	"Arc/handlers"
	"Arc/model"
	"Arc/repository"
	"Arc/repository/postgres"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := repository.Initialize()
	if err != nil {
		return
	}

	e := echo.New()

	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading environment")
	}

	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClimes{},
		SigningKey: []byte(os.Getenv("SECRET")),
	}
	addUserHandlers(e, config)
	addWebSiteHandlers(e, config)
	addPlanHandlers(e, config)
	addAdminHandler(e, config)
	e.Logger.Fatal(e.Start(":3000"))
}

func addUserHandlers(e *echo.Echo, config middleware.JWTConfig) {
	repo := &postgres.UserRepositoryImpl{
		Db: repository.Db,
	}

	h := handlers.NewUserHandler(repo)

	gp := e.Group("/users")
	gp.Use(middleware.JWTWithConfig(config))
	gp.GET("/", h.GetAllUsers)
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

func addPlanHandlers(e *echo.Echo, config middleware.JWTConfig) {
	gp := e.Group("/user/plan")
	gp.Use(middleware.JWTWithConfig(config))
	gp.POST("/", handlers.AddUserPlan)
	gp.GET("/:user_id", handlers.GetUserPlan)
	gp.PUT("/:user_id", handlers.UpdateUserPlan)
	gp.DELETE("/:user_id", handlers.DeleteUserPlan)
}

func addAdminHandler(e *echo.Echo, config middleware.JWTConfig) {
	repo := &postgres.AdminRepositoryImpl{
		Db: repository.Db,
	}

	h := handlers.NewAdminHandler(repo)
	e.POST("/admin/signup", h.Signup)
	e.POST("/admin/login", h.Login)
}
