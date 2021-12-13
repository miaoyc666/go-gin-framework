package common

/*
File name    : msg.go
Author       : miaoyc
Create date  : 2021/12/13 4:07 下午
Description  :
*/

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
