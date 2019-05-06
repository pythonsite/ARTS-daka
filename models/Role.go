package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Role struct {
	Id int `from:"Id"`
	Name string `from:"Name"`
	Seq int
	RoleResourceRel []*RoleResourceRel `orm:"reverse(many)" json:"-"`  //设置一对多的反向关系
	RoleBackendUserRel []*RoleBackendUserRel `orm:"reverse(many)" json:"-"` //设置一对多的反向关系
}

type RoleQueryParam struct {
	BaseQueryParam
	NameLike string
}

func(a *Role) TableName() string {
	return RoleTBName()
}

// RolePageList 获取分页数据
func RolePageList(params *RoleQueryParam)([]*Role, int64) {
	query := orm.NewOrm().QueryTable(RoleTBName())
	data := make([]*Role, 0)
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Seq":
		sortorder = "Seq"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("name__istartswith", params.NameLike)
	total, _ := query.Count()
	count, _ := query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	fmt.Println(count)
	return data, total
}

// RoleDataList 获取角色列表
func RoleDataList(params *RoleQueryParam) []*Role {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := RolePageList(params)
	return data
}

// RoleBatchDelete 批量删除
func RoleBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(RoleTBName())
	num, err := query.Filter("id__in",ids).Delete()
	return num, err
}

// RoleOne 获取单条
func RoleOne(id int) (*Role, error) {
	o := orm.NewOrm()
	m := Role{Id:id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}








