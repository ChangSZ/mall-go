package dto

import "time"

type SmsFlashPromotionSession struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	StartTime  string    `json:"startTime"`
	EndTime    string    `json:"endTime"`
	Status     int32     `json:"status"`
	CreateTime time.Time `json:"createTime"`
}

type SmsFlashPromotionSessionDetail struct {
	SmsFlashPromotionSession `json:",inline"`
	ProductCount             int64 `json:"productCount"` // 商品数量
}
