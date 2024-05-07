package sms_home_advertise

import "time"

// SmsHomeAdvertise 首页轮播广告表
//
//go:generate gormgen -structs SmsHomeAdvertise -input .
type SmsHomeAdvertise struct {
	Id         int64     //
	Name       string    //
	Type       int32     // 轮播位置：0->PC首页轮播；1->app首页轮播
	Pic        string    //
	StartTime  time.Time `gorm:"time"` //
	EndTime    time.Time `gorm:"time"` //
	Status     int32     // 上下线状态：0->下线；1->上线
	ClickCount int32     // 点击数
	OrderCount int32     // 下单数
	Url        string    // 链接地址
	Note       string    // 备注
	Sort       int32     // 排序
}
