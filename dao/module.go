package dao

import "time"

/*
File name    : module.go
Author       : miaoyc
Create time  : 2022/11/29 20:32
Update time  : 2022/11/29 20:32
Description  : 数据模型定义
*/

type Test struct {
	Time  time.Time
	Key   string
	Value string
}
