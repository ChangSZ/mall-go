package sms_coupon_history

import "time"

// SmsCouponHistory 优惠券使用、领取历史表
//
//go:generate gormgen -structs SmsCouponHistory -input .
type SmsCouponHistory struct {
	Id             int64     //
	CouponId       int64     //
	MemberId       int64     //
	CouponCode     string    //
	MemberNickname string    // 领取人昵称
	GetType        int32     // 获取类型：0->后台赠送；1->主动获取
	CreateTime     time.Time `gorm:"autoCreateTime"` //
	UseStatus      int32     // 使用状态：0->未使用；1->已使用；2->已过期
	UseTime        time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 使用时间
	OrderId        int64     // 订单编号
	OrderSn        string    // 订单号码
}
