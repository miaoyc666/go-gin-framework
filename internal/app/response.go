package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go-gin-example/common"
	"simple-go-gin-example/internal/e"
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

var (
	ApiKeyError     = genResponseBody(e.ApiKeyError, "", "")
	ReqParamError   = genResponseBody(e.ReqParamError, "", "")
	ReqContentError = genResponseBody(e.ReqContentError, "", "")
	ServerError     = genResponseBody(e.ServerError, "", "")
)

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  common.GetMsg(errCode),
		Data: data,
	})
	return
}

// SuccessResponse 请求成功响应结果
func (g *Gin) SuccessResponse(httpCode int, data interface{}) {
	g.C.JSON(httpCode, data)
}

// ErrorResponse 请求失败响应结果
func (g *Gin) ErrorResponse(httpCode int, data interface{}) {
	g.C.AbortWithStatusJSON(httpCode, data)
}

// DataResponse 查询数据成功响应结果
func (g *Gin) DataResponse(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	})
}

// genResponseBody 生成响应消息体
func genResponseBody(status int, data interface{}, msg string) map[string]interface{} {
	if msg == "" {
		msg = e.GetMsg(status)
	}
	return map[string]interface{}{"status": status, "data": data, "msg": msg}
}
