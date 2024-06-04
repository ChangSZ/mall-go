package dto

import "time"

type SmsFlashPromotion struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	Status     int32     `json:"status"`
	CreateTime time.Time `json:"createTime"`
}
