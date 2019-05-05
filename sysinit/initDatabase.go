package sysinit

import (
	_ "sdrms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func initDataBase() {

	dbType := beego.AppConfig.String("db_type")
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	dbName := beego.AppConfig.String(dbType + "::db_name")
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	dbCharset := beego.AppConfig.String(dbType + "::db_charset")
	_ = orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
	dbPort+")/"+dbName+"?charset="+dbCharset, 30)

	isDev := beego.AppConfig.String("runmode") == "dev"
	// 自动建表
	_ = orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}