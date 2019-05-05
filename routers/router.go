package routers

import (
	"sdrms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
    beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")

	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")

	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")
}
