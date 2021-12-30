package logger

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"time"
)

/*
File name    : logger.go
Author       : miaoyc
Create date  : 2021/12/21 12:45 上午
Description  : 日志文件相关
*/

/*
logrus内置日志级别如下：
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}
*/

var (
	logger *logrus.Logger
)

func getLog() *logrus.Logger {
	logFilePath := getLogFilePath()
	writer, err := rotateLogs.New(
		logFilePath+".%Y%m%d%H%M",
		rotateLogs.WithLinkName(logFilePath),      // 生成软链，指向最新日志文件
		rotateLogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotateLogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		panic(err.Error())
	}
	// 设置日志级别
	var level logrus.Level
	level.UnmarshalText([]byte(getLogLevel()))
	if err != nil {
		panic(err.Error())
	}
	logger = &logrus.Logger{
		Out:   writer,
		Level: level,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05+800",
			LogFormat:       "[%time%][%lvl%]%msg%\n",
		},
	}
	return logger
}

// Setup initialize the log instance
func Setup() {
	logger = getLog()
}

// Panic output logs at panic level
func Panic(v ...interface{}) {
	logger.Panic(v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Fatalf output logs at fatal level
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorf output logs at error level
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Trace output logs at trace level
func Trace(v ...interface{}) {
	logger.Trace(v...)
}
