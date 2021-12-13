package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go-gin-example/common"
)

/*
File name    : test.go
Author       : miaoyc
Create date  : 2021/12/13 4:09 下午
Description  :
*/

// GetTest
// @Summary test
// @Description test, return hello
// @Produce  json
// @Success 200 {string} string    "ok"
// @Router /test [get]
func GetTest(c *gin.Context) {
	appG := common.Gin{C: c}
	res := "hello world!"
	appG.Response(http.StatusOK, common.SUCCESS, res)
}
