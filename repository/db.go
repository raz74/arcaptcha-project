package repository

import (
	"Arc/model"
	"fmt"

	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Initialize() error {

	dsn := getCofig()
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db = database

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
		
	}

	return nil
}

func CreateAdmin(admin *model.Admin) error {
	return Db.Create(admin).Error
}

func getCofig() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
	return dsn
}
