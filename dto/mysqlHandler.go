package dto

import (
	"simple-go-gin-example/internal/module/mysqlModule"
	"simple-go-gin-example/internal/pkg/setting"
	"simple-go-gin-example/lib"
)

/*
File name    : mysqlHandler.go
Author       : miaoyc
Create Date  : 2023/7/18 17:17
Update Date  : 2023/7/18 17:17
Description  : mysql操作方法
*/

// MysqlHandler
//
//	@Description: Mysql数据操作实现
type MysqlHandler struct {
}

// InitDBConnections 初始化数据库连接
func (m *MysqlHandler) InitDBConnections() {
	maxConnections := setting.GlobalConf.Database.MaxConnections
	if maxConnections == 0 {
		maxConnections = 50
	}
	lib.InitMysqlConnection(
		setting.GlobalConf.Database.Host,
		setting.GlobalConf.Database.Port,
		setting.GlobalConf.Database.User,
		setting.GlobalConf.Database.Password,
		setting.GlobalConf.Database.Db,
		maxConnections)
	m.AutoMigrate()
}

func (m *MysqlHandler) AutoMigrate() error {

	err := lib.MysqlClient.AutoMigrate(
		&mysqlModule.Test{})
	if err != nil {
		return err
	}
	return nil
}

// TestMysqlHandler mysql数据库操作实例
var TestMysqlHandler MysqlHandler
