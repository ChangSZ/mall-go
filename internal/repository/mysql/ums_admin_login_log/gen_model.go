package ums_admin_login_log

import "time"

// UmsAdminLoginLog 后台用户登录日志表
//
//go:generate gormgen -structs UmsAdminLoginLog -input .
type UmsAdminLoginLog struct {
	Id         int64     //
	AdminId    int64     //
	CreateTime time.Time `gorm:"autoCreateTime"` //
	Ip         string    //
	Address    string    //
	UserAgent  string    // 浏览器登录类型
}
