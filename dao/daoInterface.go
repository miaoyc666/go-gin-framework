package dao

/*
File name    : daoInterface.go
Author       : miaoyc
Create time  : 2023/7/18 12:14
Update time  : 2023/7/18 12:14
Description  :
*/

// ModuleHandler
//
//	@Description: 数据库操作接口定义
type ModuleHandler interface {
	BatchGetValues(name []string) []Test
}

// dto操作接口
var module ModuleHandler

func CustomBatchGetValue(keys []string) []Test {
	return module.BatchGetValues(keys)
}

// Setup
//
//	@Description: 初始化失陷检测数据读取模块
//	@param handler DetectHandler
func Setup(handler ModuleHandler) {
	module = handler
}
