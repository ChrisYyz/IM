package service

import (
	"IM/models"
	"fmt"
	"strconv"

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
// @Param Name query string false "User Name"
// @Param Password query string false "Enter password"
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
		c.JSON(400, gin.H{
			"message": "Password mismatch!",
		})
	}
	user.Password = password
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "Create user success!",
	})
}

// DeleteUser
// @Summary Delete Exist User
// @Tags User
// @Param id query string false "Deleted ID"
// @Success 200 {string} json{"code", "message"}
// @Router /user/DeleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	ret := models.DeleteUser(user)
	if ret.Error != nil {
		c.JSON(400, gin.H{
			"message": ret.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User deleted!",
	})
}
