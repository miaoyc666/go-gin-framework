package setting

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"time"
)

/*
File name    : setting.go
Author       : miaoyc
Create date  : 2021/12/20 5:10 下午
Description  : 配置信息读取
*/

type Server struct {
	RunMode      string				`yaml:"RunMode"`
	ServerAddr   string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Swagger      bool
}

var ServerSetting = &Server{}

// Setup initialize the configuration instance
func Setup() {
	yamlFile, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Printf("read config is failed err: %s", err)
	}
	mapTo(yamlFile, ServerSetting)
}

// mapTo map section
func mapTo(fileStream []byte, v interface{}) {
	err := yaml.Unmarshal(fileStream, v)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
}
