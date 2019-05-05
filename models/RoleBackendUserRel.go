package models

import (
	"time"
)

type RoleBackendUserRel struct {
	Id int
	Role *Role `orm:"rel(fk)"` //外键
	BackendUser *BackendUser `orm:"rel(fk)"` // 外键
	Created time.Time `orm:"auto_now_add;type(datatime)"`
}

// 设置表名
func(a *RoleBackendUserRel) TableName() string {
	return RoleBackendUserRelTBName()
}

