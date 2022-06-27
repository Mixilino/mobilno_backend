package database

import (
	"backend/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DBInstance *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Cant load env string")
	}
	dsn := os.Getenv("POSTGRES")
	DBInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cant open DB")
	}
	err = DBInstance.AutoMigrate(&model.Task{})
	if err != nil {
		panic("Cant migrate Task into DB")
	}
	err = DBInstance.AutoMigrate(&model.User{})
	if err != nil {
		panic("Cant migrate User into DB")
	}
}
