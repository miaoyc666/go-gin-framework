package routers

import (
	"github.com/gin-gonic/gin"
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

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/test", v1.GetTest)

	return r
}
