package dto

/*
File name    : dtoInterface.go
Author       : miaoyc
Create Date  : 2022/11/4 16:57
Update Date  : 2022/11/4 17:15
Description  : dto对外公开暴露接口，建议其它模块只使用dto包的这些接口
*/

// DBHandler
//
//	@Description: 数据库操作接口定义
type DBHandler interface {
	InitDBConnections() // 初始化数据库连接
	AutoMigrate() error
}

// dto操作接口
var db DBHandler

// InitDBConnections 初始化数据库连接
func InitDBConnections() {
	db.InitDBConnections()
}

// AutoMigrate AutoMigrate
func AutoMigrate() error {
	return db.AutoMigrate()
}

// Setup
//
//	@Description: 初始化失陷检测数据读取模块
//	@param handler DetectHandler
func Setup(handler DBHandler) {
	db = handler
}
