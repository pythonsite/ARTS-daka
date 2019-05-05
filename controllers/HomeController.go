package controllers

import (
	"github.com/astaxie/beego/logs"
	"sdrms/models"
	"sdrms/utils"
	"strings"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	//判断是否登录
	c.checkLogin()
	c.setTpl()
}

func (c *HomeController) Page404() {
	c.setTpl()
}

func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("home/error.html", "shared/layout_pullbox.html")
}

func (c *HomeController) Login() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "home/login_footerjs.html"
	c.setTpl("home/login.html", "shared/layout_base.html")
}

func (c *HomeController) DoLogin() {
	username := strings.TrimSpace(c.GetString("UserName"))
	userpwd := strings.TrimSpace(c.GetString("UserPwd"))
	logs.Info("user:%s,pwd:%s",username,userpwd)
	if len(username) == 0 || len(userpwd) == 0 {
		c.jsonResult(utils.JRCodeFailed, "用户名或密码错误","")
	}
	userpwd = utils.String2MD5(userpwd)
	user,err := models.BackendUserOneByUserName(username, userpwd)
	if user != nil && err == nil {
		if user.Status == utils.Disabled {
			c.jsonResult(utils.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}
		//保存用户信息到session
		err = c.setBackendUser2Session(user.Id)
		if err != nil {
			logs.Error(err)
			c.jsonResult(utils.JRCodeFailed, "set session error","")
		} else {
			//获取用户信息
			c.jsonResult(utils.JRCodeSucc, "登录成功", "")
		}

 	} else {
 		c.jsonResult(utils.JRCodeFailed, "用户或者密码错误","")
	}
}

func (c *HomeController) Logout() {
	user := models.BackendUser{}
	c.SetSession("backenduser", user)
	c.pageLogin()
}

