package logger

import (
	"fmt"
	"simple-go-gin-example/internal/pkg/setting"
)

/*
File name    : file.go
Author       : miaoyc
Create date  : 2021/12/21 12:45 上午
Description  : 日志文件相关
*/

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.GlobalConf.Log.LogPath)
}

// getLogLevel get the log level
func getLogLevel() string {
	return fmt.Sprintf("%s", setting.GlobalConf.Log.LogLevel)
}
