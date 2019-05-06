package controllers

import (
	"ARTS-daka/models"
	"ARTS-daka/utils"
	"encoding/json"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Prepare() {
	// 先执行
	c.BaseController.Prepare()
	c.checkAuthor("DataGrid", "DataList", "UpdateSeq")
}

// DataGrid 角色管理首页 表格获取数据
func (c *RoleController) DataGrid() {
	var params models.RoleQueryParam
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	data, total := models.RolePageList(&params)
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//DataList 角色列表
func (c *RoleController) DataList() {
	var params = models.RoleQueryParam{}
	data := models.RoleDataList(&params)
	c.jsonResult(utils.JRCodeSucc, "",data)
}