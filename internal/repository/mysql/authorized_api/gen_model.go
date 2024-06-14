package authorized_api

import "time"

// AuthorizedApi 已授权接口地址表
//
//go:generate gormgen -structs AuthorizedApi -input .
type AuthorizedApi struct {
	Id          int64     // 主键
	BusinessKey string    // 调用方key
	Method      string    // 请求方式
	Api         string    // 请求地址
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 更新时间
	UpdatedUser string    // 更新人
}
