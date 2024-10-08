package models

import (
	"IM/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LogoutTime    *time.Time
	IsLogout      bool
	DeviceInfo    string

	IsAdmin bool
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("Name=?", name).Find(&user)
	return user
}

func FindUserByPhone(phone string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("Phone=?", phone).Find(&user)
	return user
}

func FindUserByEmail(email string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("Email=?", email).Find(&user)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Where("id=?", user.ID).Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Where("id=?", user.ID).Model(&user).Updates(UserBasic{
		Name:     user.Name,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}
