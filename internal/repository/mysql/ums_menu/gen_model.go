package ums_menu

import "time"

// UmsMenu 后台菜单表
//
//go:generate gormgen -structs UmsMenu -input .
type UmsMenu struct {
	Id         int64     //
	ParentId   int64     // 父级ID
	CreateTime time.Time `gorm:"time"` // 创建时间
	Title      string    // 菜单名称
	Level      int32     // 菜单级数
	Sort       int32     // 菜单排序
	Name       string    // 前端名称
	Icon       string    // 前端图标
	Hidden     int32     // 前端隐藏
}
