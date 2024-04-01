package ums_permission

import "time"

// UmsPermission 后台用户权限表
//
//go:generate gormgen -structs UmsPermission -input .
type UmsPermission struct {
	Id         int64     //
	Pid        int64     // 父级权限id
	Name       string    // 名称
	Value      string    // 权限值
	Icon       string    // 图标
	Type       int32     // 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）
	Uri        string    // 前端资源路径
	Status     int32     // 启用状态；0->禁用；1->启用
	CreateTime time.Time `gorm:"time"` // 创建时间
	Sort       int32     // 排序
}
