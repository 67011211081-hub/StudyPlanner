package config

import (
	"fmt"

	"studyplanner/models"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("studyplanner.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = database

	DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	)

	fmt.Println("เชื่อมฐานข้อมูลสำเร็จ")
}
