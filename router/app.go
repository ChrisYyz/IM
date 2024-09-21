package router

import (
	"IM/service"

	"IM/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/index", service.GetIndex)
	router.GET("/user/GetUserList", service.GetUserList)
	return router
}
