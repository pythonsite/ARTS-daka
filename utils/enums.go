package utils

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至网址
	JRCode401 = 401 // 未授权
)

const (
	Deleted = iota - 1
	Disabled
	Enable
)

