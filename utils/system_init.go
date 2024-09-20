package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config: App", viper.Get("app"))
	fmt.Println("config: MySQL", viper.Get("mysql"))
}

func InitMySQL() *gorm.DB {
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.DNS")), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
