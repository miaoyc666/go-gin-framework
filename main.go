package main

/*
File name    : main.go
Author       : miaoyc
Create date  : 2021/12/13 3:37 下午
Description  :
*/

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	config "github.com/spf13/viper"

	"simple-go-gin-example/routers"
)

func main() {
	// init config
	config.SetConfigType("yaml")
	config.SetConfigFile("./conf/config.yaml")
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("read config is failed err: %s", err)
	}

	// init gin server
	gin.SetMode(config.GetString("server.runMode"))
	log.Printf(config.GetString("server.runMode"))
	routersInit := routers.InitRouter()
	readTimeout := config.GetDuration("server.readTimeout") * time.Second
	writeTimeout := config.GetDuration("server.writeTimeout") * time.Second
	endPoint := config.GetString("server.addr")
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
