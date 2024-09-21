package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //singleton

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config App inited!")
}

func InitMySQL() {
	temp, err := gorm.Open(mysql.Open(viper.GetString("mysql.DNS")), &gorm.Config{})
	DB = temp
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config MySQL inited!")
}

func GetDatabase() *gorm.DB {
	return DB
}
