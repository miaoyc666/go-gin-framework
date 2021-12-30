package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

/*
File name    : setting.go
Author       : miaoyc
Create date  : 2021/12/20 5:10 下午
Description  : 配置信息读取
*/

type YamlConfig struct {
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
}

type Server struct {
	RunMode      string        `yaml:"runMode"`
	HttpPort     int32         `yaml:"httpPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	Swagger      bool          `yaml:"swagger"`
}

type Log struct {
	LogPath  string `yaml:"logPath"`
	LogLevel string `yaml:"logLevel"`
}

var (
	GlobalConf YamlConfig
)

func getConf() {
	vip := viper.New()
	vip.SetConfigType("yaml")
	vip.SetConfigFile("conf/config.yaml")
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	err := vip.Unmarshal(&GlobalConf)
	fmt.Println(GlobalConf)
	if err != nil {
		panic(err)
	}
}

// Setup initialize the configuration instance
func Setup() {
	getConf()
}
