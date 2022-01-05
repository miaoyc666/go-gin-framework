package routers

import (
	"github.com/gin-gonic/gin"
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

	r.GET("/api/v1/health", v1.HealthCheck)
	apiV1 := r.Group("/api/v1", v1.AuthMiddleware)
	apiV1.GET("/get_test", v1.GetTest)
	apiV1.POST("/post_test", v1.PostTest)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
