package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go-gin-example/internal/app"
)

/*
File name    : test.go
Author       : miaoyc
Create date  : 2021/12/13 4:09 下午
Description  :
*/

const (
	_PARAM  = "param"
	_APIKEY = "apikey"
)

// HealthCheck
// @Summary 健康检查
// @Description API接口健康检查，Get请求
// @Produce json
// @Success 10000 {string} string "ok"
// @Router /api/v1/health [get]
func HealthCheck(c *gin.Context) {
	// TODO 执行关键业务流程检查
	appG := app.Gin{C: c}
	appG.SuccessResponse(http.StatusOK, map[string]interface{}{"status": "ok"})
}

// GetTest
// @Summary test
// @Description test, return hello
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/get_test [get]
// @Param apikey query string true "apikey"
func GetTest(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.DataResponse(map[string]interface{}{"hello": "world"})
}

// PostTest
// @Summary test
// @Description test, return hello
// @accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/post_test [post]
// @Param apikey body string true "apikey"
// @Param param body string true "param"
func PostTest(c *gin.Context) {
	refApikey := c.GetString(_APIKEY)
	refParam := c.GetString(_PARAM)
	appG := app.Gin{C: c}
	appG.DataResponse(map[string]interface{}{refApikey: refParam})
}

func AuthMiddleware(c *gin.Context) {
	// 1.获取接口请求参数
	appG := app.Gin{C: c}
	var reqParam ReqParam
	if c.Request.Method == "GET" {
		err := c.ShouldBindQuery(&reqParam)
		if err != nil {
			appG.ErrorResponse(http.StatusOK, app.ReqParamError)
			return
		}
	} else {
		contentType := c.Request.Header.Get("Content-Type")
		if contentType != "application/json" {
			appG.ErrorResponse(http.StatusOK, app.ReqContentError)
			return
		}
		if c.ShouldBind(&reqParam) != nil {
			appG.ErrorResponse(http.StatusOK, app.ReqParamError)
			return
		}
	}
	// 2.验证apikey有效性
	isApikeyValid, err := checkApiKey(reqParam.Apikey)
	if err != nil || !isApikeyValid {
		appG.ErrorResponse(http.StatusOK, app.ApiKeyError)
		return
	}

	// 3.传参给Context
	c.Set(_APIKEY, reqParam.Apikey)
	c.Set(_PARAM, reqParam.Param)

	// c.Next()语句之前的处理为前置拦截
	c.Next()
	// c.Next()语句之后的处理为后置拦截

}

// checkApiKey 检查apikey有效性，此处为测试，实际使用添加业务实现即可
func checkApiKey(apiKey string) (bool, error) {
	if apiKey != "" {
		return true, nil
	}
	return false, nil
}
