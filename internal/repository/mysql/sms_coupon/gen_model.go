package sms_coupon

import "time"

// SmsCoupon 优惠券表
//
//go:generate gormgen -structs SmsCoupon -input .
type SmsCoupon struct {
	Id           int64     //
	Type         int32     // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
	Name         string    //
	Platform     int32     // 使用平台：0->全部；1->移动；2->PC
	Count        int32     // 数量
	Amount       float64   // 金额
	PerLimit     int32     // 每人限领张数
	MinPoint     float64   // 使用门槛；0表示无门槛
	StartTime    time.Time `gorm:"time"` //
	EndTime      time.Time `gorm:"time"` //
	UseType      int32     // 使用类型：0->全场通用；1->指定分类；2->指定商品
	Note         string    // 备注
	PublishCount int32     // 发行数量
	UseCount     int32     // 已使用数量
	ReceiveCount int32     // 领取数量
	EnableTime   time.Time `gorm:"time"` // 可以领取的日期
	Code         string    // 优惠码
	MemberLevel  int32     // 可领取的会员类型：0->无限时
}
