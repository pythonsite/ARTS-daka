package controllers

import (
	"ARTS-daka/models"
	"ARTS-daka/utils"
	"strings"
)

type ResourceController struct {
	BaseController
}

func (c *ResourceController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("TreeGrid", "UserMenuTree", "ParentTreeGrid", "Select")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

func (c *ResourceController) UrlFor2LinkOne(urlfor string) string {
	if len(urlfor) == 0 {
		return ""
	}
	strs := strings.Split(urlfor, ",")
	if len(strs) == 1 {
		return c.URLFor(strs[0])
	} else if len(strs) > 1 {
		var values []interface{}
		for _, val := range strs[1:] {
			values = append(values, val)
		}
		return c.URLFor(strs[0], values)
	}
	return ""
}

func (c *ResourceController) UrlFor2Link(src []*models.Resource) {
	for _, item := range src{
		item.LinkUrl = c.UrlFor2LinkOne(item.UrlFor)
	}
}

func (c *ResourceController) UserMenuTree() {
	userid := c.curUser.Id
	// 获取用户权限管理菜单列表
	tree := models.ResourceTreeGridByUserId(userid, 1)
	c.UrlFor2Link(tree)
	c.jsonResult(utils.JRCodeSucc, "", tree)
}

func (c *ResourceController) Index() {
	//需要权限控制
	c.checkAuthor()
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "resource/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "resource/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("ResourceController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("ResourceController", "Delete")
}

//CheckUrlFor 填写urlFor时进行验证
func (c *ResourceController) CheckUrlFor() {
	urlfor := c.GetString("urlfor")
	link := c.UrlFor2LinkOne(urlfor)
	if len(link) > 0 {
		c.jsonResult(utils.JRCodeSucc, "解析成功", link)
	} else {
		c.jsonResult(utils.JRCodeFailed,"解析失败",link)
	}
}


