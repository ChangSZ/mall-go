package sms_flash_promotion_session

import "time"

// SmsFlashPromotionSession 限时购场次表
//
//go:generate gormgen -structs SmsFlashPromotionSession -input .
type SmsFlashPromotionSession struct {
	Id         int64     // 编号
	Name       string    // 场次名称
	StartTime  time.Time `gorm:"time"` // 每日开始时间
	EndTime    time.Time `gorm:"time"` // 每日结束时间
	Status     int32     // 启用状态：0->不启用；1->启用
	CreateTime time.Time `gorm:"autoCreateTime"` // 创建时间
}
