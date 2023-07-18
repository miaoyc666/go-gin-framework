package dao

/*
File name    : mysqlHandler.go
Author       : miaoyc
Create Date  : 2023/7/18 12:41
Update Date  : 2023/7/18 12:41
Description  : mysql操作方法
*/

import (
	"simple-go-gin-example/internal/logger"
	"simple-go-gin-example/internal/module/mysqlModule"
	"simple-go-gin-example/lib"
)

var (
	tableModelMap map[string]interface{} // 表名称与model的映射关系
	testModel     mysqlModule.Test
)

func init() {
	initIndexMap()
}

func initIndexMap() {
	tableModelMap = make(map[string]interface{})
	tableModelMap[mysqlModule.TableTest] = testModel
}

// MysqlHandler
//
//	@Description: Mysql数据操作实现
type MysqlHandler struct {
}

// BatchGetValues 批量获取values
func (m *MysqlHandler) BatchGetValues(keys []string) (infoList []Test) {
	var results []mysqlModule.Test
	tx := lib.MysqlClient.Where("key IN ?", keys).Find(&results)
	if tx.Error != nil {
		logger.Errorf("BatchGetValues query db error: %s", tx.Error)
		return infoList
	}
	// 转换数据格式
	for _, v := range results {
		p := mysqlTest2Test(v)
		infoList = append(infoList, p)
	}
	return infoList
}

// TestMysqlHandler Mysql数据库操作实例
var TestMysqlHandler MysqlHandler
