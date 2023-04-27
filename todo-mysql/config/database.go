package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"todo-mysql/models"
)


var Database *gorm.DB
var err error
//var url = ""+connect.Username+":"+connect.Password+"@tcp("+connect.Host+":"+strconv.Itoa(connect.Port)+")/"+connect.Database+""

func Connect(connect models.Connect) {
	var url = ""+connect.Username+":"+connect.Password+"@tcp("+connect.Host+":"+strconv.Itoa(connect.Port)+")/"+connect.Database+""
	fmt.Println("URL", url)
	Database, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection Refused!")
	}
	Database.AutoMigrate(&models.Todo{})
}