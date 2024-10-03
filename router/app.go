package router

import (
	"IM/service"

	"IM/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	router := gin.Default() // Logger and Recovery middleware already attached
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // router.Handle("GET"...), regiser handler

	router.GET("/user/GetUserList", service.GetUserList)
	router.GET("/user/CreateUser", service.CreateUser)
	router.GET("/user/DeleteUser", service.DeleteUser)
	router.POST("/user/UpdateUser", service.UpdateUser)
	return router
}
