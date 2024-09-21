package service

import (
	"IM/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags HomePage
// @Success 200 {string} json{"code", "message"}
// @Router /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}
