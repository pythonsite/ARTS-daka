package models

import "ARTS-daka/utils"

type JsonResult struct {
	Code utils.JsonResultCode `json:"code"`
	Msg string `json:"msg"`
	Obj interface{} `json:"obj"`
}

//BaseQueryParam 用于查询的类
type BaseQueryParam struct {
	Sort string `json:"sort"`
	Order string `json:"order"`
	Offset int64 `json:"offset"`
	Limit int `json:"limit"`
}
