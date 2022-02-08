package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go-gin-example/internal/app"
	"simple-go-gin-example/internal/logger"
	"time"
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
	// 记录请求起始时间
	t := time.Now()

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

	// 获取response处理
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	// c.Next()语句之前的处理为前置拦截
	c.Next()
	// c.Next()语句之后的处理为后置拦截
	// 记录业务逻辑结束时间
	t2 := time.Since(t)
	printMatchLog(c, reqParam, blw.body, t2)
}

// checkApiKey 检查apikey有效性，此处为测试，实际使用添加业务实现即可
func checkApiKey(apiKey string) (bool, error) {
	if apiKey != "" {
		return true, nil
	}
	return false, nil
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func printMatchLog(c *gin.Context, reqParam ReqParam, data interface{}, cost time.Duration) {
	// match log 格式如下：
	/*
		[2021-12-22 21:28:39+0800][INFO][ip][apikey][url][POST][{"apikey":"xxxx","param":"1.2.3.4"}][query_response=[]; cost=0.0364110469818]
	*/
	queryRequest, _ := json.Marshal(reqParam)
	msg := fmt.Sprintf("[%s][%s][%s][%s][%s][query_response=%s; cost=%s]",
		c.ClientIP(), reqParam.Apikey, c.Request.URL.Path, c.Request.Method, string(queryRequest), data, cost)
	logger.Info(msg)
}
