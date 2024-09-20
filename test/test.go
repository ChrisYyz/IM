package main

import (
	"IM/models"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate((&models.UserBasic{}))
	user := &models.UserBasic{}
	user.Name = "yyz"
	db.Create(user)

	//Read back
	fmt.Println(db.First(user, 1))

	//Update db
	db.Model(user).Update("password", "1234")
}
