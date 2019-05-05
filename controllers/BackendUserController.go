package controllers

import (
	"ARTS-daka/models"
	"encoding/json"
)

type BackendUserController  struct {
	BaseController
}

func (c *BackendUserController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("dataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *BackendUserController) Index() {
	// 是否限制更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backenduser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backenduser/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("BackendUserController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("BackendUserController", "Delete")
}

func (c *BackendUserController) DataGrid() {
	var params models.BackendUserQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	data, total := models.BackendUserPageList(&params)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
