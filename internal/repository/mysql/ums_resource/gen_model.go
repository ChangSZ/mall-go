package ums_resource

import "time"

// UmsResource 后台资源表
//
//go:generate gormgen -structs UmsResource -input .
type UmsResource struct {
	Id          int64     //
	CreateTime  time.Time `gorm:"time"` // 创建时间
	Name        string    // 资源名称
	Url         string    // 资源URL
	Description string    // 描述
	CategoryId  int64     // 资源分类ID
}
