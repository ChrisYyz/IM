package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // print to console
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	temp, err := gorm.Open(mysql.Open(viper.GetString("mysql.DNS")), &gorm.Config{Logger: dbLogger})
	DB = temp
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config MySQL inited!")
}

func GetDatabase() *gorm.DB {
	return DB
}
