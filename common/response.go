package common

import (
	"github.com/gin-gonic/gin"
)

/*
File name    : response.go
Author       : miaoyc
Create date  : 2021/12/13 4:02 下午
Description  :
*/

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}
