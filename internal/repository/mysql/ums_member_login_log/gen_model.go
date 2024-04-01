package ums_member_login_log

import "time"

// UmsMemberLoginLog 会员登录记录
//
//go:generate gormgen -structs UmsMemberLoginLog -input .
type UmsMemberLoginLog struct {
	Id         int64     //
	MemberId   int64     //
	CreateTime time.Time `gorm:"autoCreateTime"` //
	Ip         string    //
	City       string    //
	LoginType  int32     // 登录类型：0->PC；1->android;2->ios;3->小程序
	Province   string    //
}
