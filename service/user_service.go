package service

import (
	"IM/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary All Users
// @Tags User
// @Success 200 {string} json{"code", "message"}
// @Router /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary Add New User
// @Tags User
// @Param name query string false "User Name"
// @Param password query string false "Enter password"
// @Param rePassword query string false "Enter password again"
// @Success 200 {string} json{"code", "message"}
// @Router /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("Name")
	password := c.Query("Password")
	rePassword := c.Query("rePassword")
	fmt.Println(password + " | " + rePassword)
	if password != rePassword {
		c.JSON(-1, gin.H{
			"message": "Password mismatch!",
		})
	}
	user.Password = password
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "Create user success!",
	})
}
