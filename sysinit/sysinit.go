package sysinit

import (
	"github.com/astaxie/beego"
	"ARTS-daka/utils"
)

func init() {
	// 启动session
	beego.BConfig.WebConfig.Session.SessionOn = true
	// 初始化日志
	utils.InitLogs()
	// 初始化缓存
	utils.InitCache()
	// 初始化数据库
	initDataBase()
}
