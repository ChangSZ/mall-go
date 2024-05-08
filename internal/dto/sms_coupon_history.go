package dto

import "time"

type SmsCouponHistory struct {
	Id             int64     `json:"id"`
	CouponId       int64     `json:"couponId"`
	MemberId       int64     `json:"memberId"`
	CouponCode     string    `json:"couponCode"`
	MemberNickname string    `json:"memberNickname"`
	GetType        int32     `json:"getType"`
	CreateTime     time.Time `json:"createTime"`
	UseStatus      int32     `json:"useStatus"`
	UseTime        time.Time `json:"useTime"`
	OrderId        int64     `json:"orderId"`
	OrderSn        string    `json:"orderSn"`
}
