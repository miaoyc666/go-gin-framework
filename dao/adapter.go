package dao

import (
	"simple-go-gin-example/internal/module/mysqlModule"
)

/*
File name    : adapter.go
Author       : miaoyc
Create time  : 2023/7/18 13:12
Update time  : 2023/7/18 13:12
Description  : 结构体转换
*/

// mysqlTest2Test 将mysqlModule.Test转换为Test
func mysqlTest2Test(test mysqlModule.Test) (m Test) {
	m.Time = test.Time
	m.Key = test.Key
	m.Value = test.Value
	return m
}
