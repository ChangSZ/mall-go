package admin_menu

import "time"

// AdminMenu 管理员菜单栏表
//
//go:generate gormgen -structs AdminMenu -input .
type AdminMenu struct {
	Id          int64     // 主键
	AdminId     int64     // 管理员ID
	MenuId      int64     // 菜单栏ID
	CreatedAt   time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 创建时间
	CreatedUser string    // 创建人
}
