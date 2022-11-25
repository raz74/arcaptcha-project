package main

import (
	"Arc/handelers"
	"Arc/repository"

	"github.com/labstack/echo"

)

func main() {
	repository.Initialize()
	e := echo.New()
	e.POST("/signup", handelers.Signup)
	e.POST("/login", handelers.Login)
// get /user -> list of all users
	e.GET("/users", handelers.GetAllUsers)
// get /user/{id} -> get that user
	e.GET("/users/:id", handelers.GetUser)
// put /user/{id} -> put that user
	e.PUT("/users/:id", handelers.UpdateUser)
// delete /user/{id} -> delete that user 
	e.DELETE("/users/:id", handelers.DeleteUser)
	e.Logger.Fatal(e.Start(":3000"))
}
