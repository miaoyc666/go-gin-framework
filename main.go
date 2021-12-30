package main

/*
File name    : main.go
Author       : miaoyc
Create date  : 2021/12/13 3:37 下午
Description  :
*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-go-gin-example/internal/logger"
	"simple-go-gin-example/internal/pkg/setting"
	"simple-go-gin-example/routers"
	"time"
)

func init() {
	setting.Setup()
	logger.Setup()
}

func main() {
	// init gin server
	gin.SetMode(setting.GlobalConf.Server.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.GlobalConf.Server.ReadTimeout * time.Second
	writeTimeout := setting.GlobalConf.Server.WriteTimeout * time.Second
	endPoint := fmt.Sprintf(":%d", setting.GlobalConf.Server.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
