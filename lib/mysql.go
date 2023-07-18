package lib

/*
File name    : mysql.go
Author       : miaoyc
Create Date  : 2023/7/17 23:03
Update Date  : 2023/7/17 23:03
Description  : mysql数据库连接
*/

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	frameworkLogger "simple-go-gin-example/internal/logger"
)

var (
	MysqlClient *gorm.DB
)

func InitMysqlConnection(host, port, user, password, dbname string, connections int) (err error) {
	// gorm 连接mysql的数据库
	MysqlClient, err = getMysqlClient(host, port, user, password, dbname) // 连接本地db
	if err != nil {
		return err
	}
	sqlDB, err := MysqlClient.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(connections) //设置最大连接数
	sqlDB.SetMaxOpenConns(connections) //设置最大的空闲连接数
	return nil
}

func getMysqlClient(host, port, user, password, dbname string) (*gorm.DB, error) {
	// 获取postgres客户端
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, password, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		frameworkLogger.Info("Connect mysql success!")
	}
	// 关闭gorm的默认打印日志
	db.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	return db, nil
}
