package main

/*
File name    : main.go
Author       : miaoyc
Create date  : 2021/12/13 3:37 下午
Description  :
*/

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"simple-go-gin-example/dao"
	"simple-go-gin-example/dto"
	"time"

	"github.com/gin-gonic/gin"

	"simple-go-gin-example/internal/logger"
	"simple-go-gin-example/internal/pkg/setting"
	"simple-go-gin-example/routers"
)

func init() {
	var configFilePath string
	flag.StringVar(&configFilePath, "c", "", "config file path")
	flag.Parse()
	if configFilePath == "" {
		configFilePath = "./conf/config.yaml"
	}

	// init setting, logger
	setting.Setup(configFilePath)
	logger.Setup()

	// init dto and dao handler
	if setting.GlobalConf.Database.DbType == "None" {
		return
	}
	if setting.GlobalConf.Database.DbType == "mysql" {
		dao.Setup(&dao.TestMysqlHandler)
		dto.Setup(&dto.TestMysqlHandler)
	}
	// init db and auto migrate
	dto.InitDBConnections()
	dto.AutoMigrate()
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
