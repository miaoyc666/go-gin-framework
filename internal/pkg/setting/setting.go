package setting

import (
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
	Server   Server   `yaml:"server"`
	Log      Log      `yaml:"log"`
	Database Database `yaml:"database"`
}

type Server struct {
	RunMode      string        `yaml:"runMode"`
	HttpPort     int32         `yaml:"httpPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	Swagger      bool          `yaml:"swagger"`
}

type Database struct {
	DbType         string `yaml:"dbType"`
	Host           string `yaml:"host"`
	Port           string `yaml:"httpPort"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Db             string `yaml:"db"`
	MaxConnections int    `yaml:"maxConnections"`
}

type Log struct {
	LogPath  string `yaml:"logPath"`
	LogLevel string `yaml:"logLevel"`
}

var (
	GlobalConf YamlConfig
)

func getConf(configFile string) {
	vip := viper.New()
	vip.SetConfigType("yaml")
	vip.SetConfigFile(configFile)
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	err := vip.Unmarshal(&GlobalConf)
	if err != nil {
		panic(err)
	}
}

// Setup initialize the configuration instance
func Setup(configFile string) {
	getConf(configFile)
}
