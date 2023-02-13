package e

var MsgFlags = map[uint]string{
	Success:       "ok",
	Error:         "faile",
	InvalidParams: "请求的参数错误",
}

// GetMsg，获取状态码对应的信息
func GetMsg(code uint) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[Error]
}
