package controllers

import (
	"ARTS-daka/models"
	"ARTS-daka/utils"
	"github.com/astaxie/beego/orm"
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

//TreeGrid 获取所有资源的列表
func (c *ResourceController) TreeGrid() {
	tree := models.ResourceTreeGrid()
	c.UrlFor2Link(tree)
	c.jsonResult(utils.JRCodeSucc,"",tree)
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

//Edit 资源编辑页面
func (c *ResourceController) Edit() {
	// 需要权限控制
	c.checkAuthor()
	if c.Ctx.Request.Method == "POST"{
		//c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &models.Resource{}
	var err error
	if Id == 0 {
		m.Seq = 100
	} else {
		m, err = models.ResourceOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新重试")
		}
	}
	if m.Parent != nil {
		c.Data["parent"] = m.Parent.Id
	} else {
		c.Data["pareent"] = 0
	}
	// 获取可以成为当前节点的父节点的列表
	c.Data["parents"] = models.ResourceTreeGrid4Parent(Id)
	// 转换地址
	m.LinkUrl = c.UrlFor2LinkOne(m.UrlFor)
	c.Data["m"] = m
	if m.Parent != nil {
		c.Data["parent"] = m.Parent.Id
	} else {
		c.Data["parent"] = 0
	}

	c.setTpl("resource/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "resource/edit_footerjs.html"

}

//Save 资源添加编辑 保存
func (c *ResourceController) Save() {
	var err error
	o := orm.NewOrm()
	parent := &models.Resource{}
	m := models.Resource{}
	parentId, _ := c.GetInt("Parent", 0)
	//获取form中的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(utils.JRCodeFailed, "获取数据失败", m.Id)
	}
	//获取父节点
	if parentId > 0 {
		parent, err = models.ResourceOne(parentId)
		if err == nil && parent != nil {
			m.Parent = parent
		} else {
			c.jsonResult(utils.JRCodeFailed, "父节点无效", "")
		}
	}
	if m.Id == 0 {
		if _, err = o.Insert(&m); err == nil {
			c.jsonResult(utils.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.jsonResult(utils.JRCodeFailed, "添加失败", m.Id)
		}

	} else {
		if _, err = o.Update(&m); err == nil {
			c.jsonResult(utils.JRCodeSucc, "编辑成功", m.Id)
		} else {
			c.jsonResult(utils.JRCodeFailed, "编辑失败", m.Id)
		}

	}
}
