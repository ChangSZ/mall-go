package ums_resource_category

import "time"

// UmsResourceCategory 资源分类表
//
//go:generate gormgen -structs UmsResourceCategory -input .
type UmsResourceCategory struct {
	Id         int64     //
	CreateTime time.Time `gorm:"time"` // 创建时间
	Name       string    // 分类名称
	Sort       int32     // 排序
}
