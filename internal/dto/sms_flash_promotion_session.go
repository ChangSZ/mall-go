package dto

import "time"

type SmsFlashPromotionSession struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	Status     int32     `json:"status"`
	CreateTime time.Time `json:"createTime"`
}

type SmsFlashPromotionSessionDetail struct {
	SmsFlashPromotionSession `json:",inline"`
	ProductCount             int64 `json:"productCount"` // 商品数量
}
