package repository

import (
	"Arc/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var Db *gorm.DB

func Initialize() error {
	dsn := "host=localhost user=admin password=123456 dbname=postgres port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db = database

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
		// return err
	}
	Db.AutoMigrate(&model.User{})
	return nil
}


// create user
func CreateUser(user *model.User) error {
	result := Db.Create(user)
	fmt.Printf("%+v\n", result)
	return nil
}