package dto

import "time"

type SmsHomeAdvertise struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Type       int32     `json:"type"`
	Pic        string    `json:"pic"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	Status     int32     `json:"status"`
	ClickCount int32     `json:"clickCount"`
	OrderCount int32     `json:"orderCount"`
	Url        string    `json:"url"`
	Note       string    `json:"note"`
	Sort       int32     `json:"sort"`
}
