package service

import (
	"IM/models"
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary All Users
// @Tags User
// @Success 200 {string} json{"code", "message"}
// @Router /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	// data := make([]*models.UserBasic, 10)
	data := models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary Add New User
// @Tags User
// @Param Name query string false "Name"
// @Param Phone query string false "Mobile number"
// @Param Email query string false "Email"
// @Param Password query string false "Enter password"
// @Param rePassword query string false "Enter password again"
// @Success 200 {string} json{"code", "message"}
// @Router /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("Name")
	user.Email = c.Query("Email")
	user.Phone = c.Query("Phone")
	password := c.Query("Password")
	rePassword := c.Query("rePassword")

	checkUserExist := models.UserBasic{}
	checkUserExist = models.FindUserByName(user.Name)
	if checkUserExist.Name != "" {
		c.JSON(-1, gin.H{
			"message": "User Name existed!",
		})
		return
	}
	fmt.Println(user.Email)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	if password != rePassword {
		c.JSON(400, gin.H{
			"message": "Password mismatch!",
		})
		return
	}
	user.Password = password
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "User created!",
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

// UpdateUser
// @Summary Update User Information
// @Tags User
// @Param id formData string false "id"
// @Param Name formData string false "Name"
// @Param Password formData string false "Password"
// @Param Phone formData string false "Phone"
// @Param Email formData string false "Email"
// @Success 200 {string} json{"code", "message"}
// @Router /user/UpdateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("Name")
	user.Password = c.PostForm("Password")
	user.Phone = c.PostForm("Phone")
	user.Email = c.PostForm("Email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	ret := models.UpdateUser(user)
	if ret.Error != nil {
		c.JSON(400, gin.H{
			"message": ret.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User updated!",
	})
}
