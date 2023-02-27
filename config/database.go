package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo-mysql/models"
)


var Database *gorm.DB
var err error
var url = "root:root@tcp(10.100.101.48:3316)/todolist"

func Connect() {
	Database, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
		fmt.Println("Connection Refused!")
	}
	Database.AutoMigrate(&models.Todo{})
}