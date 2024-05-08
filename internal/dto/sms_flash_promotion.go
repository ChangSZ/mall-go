package dto

import "time"

type SmsFlashPromotion struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	StartDate  string    `json:"startDate"`
	EndDate    string    `json:"endDate"`
	Status     int32     `json:"status"`
	CreateTime time.Time `json:"createTime"`
}
