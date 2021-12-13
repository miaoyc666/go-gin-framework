package routers

import (
	"github.com/gin-gonic/gin"
	config "github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"simple-go-gin-example/docs"
	v1 "simple-go-gin-example/routers/api"
)

/*
File name    : routers.go
Author       : miaoyc
Create date  : 2021/12/13 3:49 下午
Description  :
*/

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api/v1"

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/test", v1.GetTest)
	}
	if config.GetBool("server.swagger") {
		apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return r
}
