package main

import (
	"IM/models"
	"IM/router"
	"IM/utils"
	"fmt"
)

func main() {
	utils.InitConfig()
	db := utils.InitMySQL()
	user := models.UserBasic{}
	db.Find(&user)
	fmt.Println(user)
	router := router.Router()
	router.Run()
}
