package sms_flash_promotion_log

import "time"

// SmsFlashPromotionLog 限时购通知记录
//
//go:generate gormgen -structs SmsFlashPromotionLog -input .
type SmsFlashPromotionLog struct {
	Id            int64     //
	MemberId      int64     //
	ProductId     int64     //
	MemberPhone   string    //
	ProductName   string    //
	SubscribeTime time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 会员订阅时间
	SendTime      time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` //
}
