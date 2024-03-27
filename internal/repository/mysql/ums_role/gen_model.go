package ums_role

import "time"

// UmsRole 后台用户角色表
//
//go:generate gormgen -structs UmsRole -input .
type UmsRole struct {
	Id          int64     //
	Name        string    // 名称
	Description string    // 描述
	AdminCount  int32     // 后台用户数量
	CreateTime  time.Time `gorm:"time"` // 创建时间
	Status      int32     // 启用状态：0->禁用；1->启用
	Sort        int32     //
}
