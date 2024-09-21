package main

import (
	"IM/router"
	"IM/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	router := router.Router()
	router.Run()
}
