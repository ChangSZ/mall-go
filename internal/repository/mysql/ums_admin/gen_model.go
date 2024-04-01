package ums_admin

import "time"

// UmsAdmin 后台用户表
//
//go:generate gormgen -structs UmsAdmin -input .
type UmsAdmin struct {
	Id         int64     //
	Username   string    //
	Password   string    //
	Icon       string    // 头像
	Email      string    // 邮箱
	NickName   string    // 昵称
	Note       string    // 备注信息
	CreateTime time.Time `gorm:"autoCreateTime"` // 创建时间
	LoginTime  time.Time `gorm:"time"`           // 最后登录时间
	Status     int32     // 帐号启用状态：0->禁用；1->启用
}
